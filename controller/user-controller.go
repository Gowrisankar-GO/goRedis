package controller

import (
	"context"
	"encoding/json"
	"go_redis/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// function used to fetch user details
func FetchUser(c *gin.Context) {

	var resp models.Response

	userId := c.Param("id")

	if userId == "" {

		resp = models.Response{Message: "userid is required", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	key := strings.ReplaceAll(models.UserKey, "id", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	userDetail, err := Defs.GetUserDetail(key)

	if err != nil {

		resp = models.Response{Message: "unable to fetch user detail", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	resp = models.Response{Message: "successfully fetched user detail", Status: 1, Data: gin.H{"userDetail": userDetail}}

	c.JSON(http.StatusOK, resp)
}

// function used to create a new user
func CreateUser(c *gin.Context) {

	var resp models.Response

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {

		resp = models.Response{Message: "unable to read request body", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	var userData models.User

	err = json.Unmarshal(bodyBytes, &userData)

	if err != nil {

		resp = models.Response{Message: "unable to read user data", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	userId, err := Defs.IncrementUserId()

	if err != nil {

		resp = models.Response{Message: "unable to create a new user", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	userData.Id = uint(userId)

	key := strings.ReplaceAll(models.UserKey, "id", strconv.Itoa(int(userId)))

	ctxone, cancelone := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defsone := models.BasicDefs{DbConn: models.Rdb, Ctx: ctxone, CtxCancel: cancelone}

	err = Defsone.SetUser(key, userData)

	if err != nil {

		resp = models.Response{Message: "unable to create a new user", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	resp = models.Response{Message: "successfully created a new user", Status: 1, Data: gin.H{"userDetail": userData}}

	c.JSON(200, resp)
}

// function used to update the details of the user
func UpdateUser(c *gin.Context) {

	var resp models.Response

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {

		resp = models.Response{Message: "unable to read request body", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	var userData models.User

	err = json.Unmarshal(bodyBytes, &userData)

	if err != nil {

		resp = models.Response{Message: "unable to read user data", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	userId := int(userData.Id)

	key := strings.ReplaceAll(models.UserKey, "id", strconv.Itoa(userId))

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	err = Defs.SetUser(key, userData)

	if err != nil {

		resp = models.Response{Message: "unable to update the user", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	resp = models.Response{Message: "successfully updated the user", Status: 1, Data: gin.H{"userDetail": userData}}

	c.JSON(http.StatusOK, resp)
}

// function used to delete the user
func DeleteUser(c *gin.Context) {

	var resp models.Response

	userId := c.Param("id")

	if userId == "" {

		resp = models.Response{Message: "userid is required", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	key := strings.ReplaceAll(models.UserKey, "id", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	err := Defs.DeleteUserTransaction(key)

	if err != nil {

		resp = models.Response{Message: "unable to delete the user", Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	resp = models.Response{Message: "successfully deleted the user", Status: 1, Data: gin.H{"userId": userId}}

	c.JSON(http.StatusOK, resp)
}
