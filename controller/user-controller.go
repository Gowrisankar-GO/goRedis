// Package controller contains handler functions used in the application
package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"redis_user_management/info"
	"redis_user_management/models"
	"redis_user_management/validator"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Function FetchUser used to fetch user details
// FetchUser godoc
// @Summary Fetch a user by ID
// @Description Retrieve a user by their unique ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/{id} [get]
func FetchUser(c *gin.Context) {

	var resp models.Response

	userId := c.Param("id")

	if userId == "" {

		InfoLog.Printf("%v",info.MsgReqUid)

		resp = models.Response{Message: info.MsgReqUid, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	key := strings.ReplaceAll(models.UserKey, "id", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	userDetail, err := Defs.GetUserDetail(key)

	if err != nil {

		InfoLog.Printf("%v: %v",info.MsgFetchUser,userId)

		resp = models.Response{Message: info.MsgFetchUser, Status: 0, Data: gin.H{}}

		if err == redis.Nil{

			c.AbortWithStatusJSON(http.StatusBadRequest, resp)

			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	InfoLog.Printf("%v: %v",info.MsgGetUser,userId)

	resp = models.Response{Message: info.MsgGetUser, Status: 1, Data: gin.H{"userDetail": userDetail}}

	c.JSON(http.StatusOK, resp)
}

// Function CreateUser used to create a new user
// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "User Information"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/create/{id} [post]
func CreateUser(c *gin.Context) {

	var resp models.Response

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		
		InfoLog.Printf("%v",info.MsgReadBody)

		resp = models.Response{Message: info.MsgReadBody, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	var userData models.User

	err = json.Unmarshal(bodyBytes, &userData)

	if err != nil {

		InfoLog.Printf("%v",info.MsgReadJson)

		resp = models.Response{Message: info.MsgReadJson, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	if err := validator.ValidateStruct(userData); err != nil {

		InfoLog.Printf("%v",info.MsgJsonValid)

		resp = models.Response{Message: info.MsgJsonValid, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return

	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	userId, err := Defs.IncrementUserId()

	if err != nil {

		ErrLog.Printf("%v",info.ErrCreateUser)

		resp = models.Response{Message: info.MsgCreateUser, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	userData.Id = uint(userId)

	key := strings.ReplaceAll(models.UserKey, "id", strconv.Itoa(int(userId)))

	ctxone, cancelone := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defsone := models.BasicDefs{DbConn: models.Rdb, Ctx: ctxone, CtxCancel: cancelone}

	err = Defsone.SetUser(key, userData)

	if err != nil {

		ErrLog.Printf("%v",info.ErrCreateUser)

		resp = models.Response{Message: info.MsgCreateUser, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	InfoLog.Printf("%v: %v",info.MsgNewUser,userId)

	resp = models.Response{Message: info.MsgNewUser, Status: 1, Data: gin.H{"userDetail": userData}}

	c.JSON(200, resp)
}

// Function UpdateUser used to update the details of the user
// UpdateUser godoc
// @Summary Update an existing user
// @Description Update the user details for the specified ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "Updated User Information"
// @Param   id  path  string  true  "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/update [put]
func UpdateUser(c *gin.Context) {

	var resp models.Response

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {

		InfoLog.Printf("%v",info.MsgReadBody)

		resp = models.Response{Message: info.MsgReadBody, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	var userData models.User

	err = json.Unmarshal(bodyBytes, &userData)

	if err != nil {

		InfoLog.Printf("%v",info.MsgReadJson)

		resp = models.Response{Message: info.MsgReadJson, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	if err := validator.ValidateStruct(userData); err != nil {

		InfoLog.Printf("%v",info.MsgJsonValid)

		resp = models.Response{Message: info.MsgJsonValid, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return

	}

	userId := int(userData.Id)

	key := strings.ReplaceAll(models.UserKey, "id", strconv.Itoa(userId))

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	err = Defs.SetUser(key, userData)

	if err != nil {

		ErrLog.Printf("%v: %v",info.ErrUpdateUser,userId)

		resp = models.Response{Message: info.MsgUpdateUser, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	InfoLog.Printf("%v: %v",info.MsgUpdatedUser,userId)

	resp = models.Response{Message: info.MsgUpdatedUser, Status: 1, Data: gin.H{"userDetail": userData}}

	c.JSON(http.StatusOK, resp)
}

// Function DeleteUser used to delete the user
// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete the user with the specified ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/delete/{id} [delete]
func DeleteUser(c *gin.Context) {

	var resp models.Response

	userId := c.Param("id")

	if userId == "" {

		InfoLog.Printf("%v",info.MsgReqUid)

		resp = models.Response{Message: info.MsgReqUid, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)

		return
	}

	key := strings.ReplaceAll(models.UserKey, "id", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	Defs := models.BasicDefs{DbConn: models.Rdb, Ctx: ctx, CtxCancel: cancel}

	err := Defs.DeleteUserTransaction(key)

	if err != nil {

		ErrLog.Printf("%V: %v",info.ErrDelUser,userId)

		resp = models.Response{Message: info.MsgDelUser, Status: 0, Data: gin.H{}}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)

		return
	}

	InfoLog.Printf("%v: %v",info.MsgDeleteUser,userId)

	resp = models.Response{Message: info.MsgDelUser, Status: 1, Data: gin.H{"userId": userId}}

	c.JSON(http.StatusOK, resp)
}
