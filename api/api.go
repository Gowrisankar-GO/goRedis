// Package api provides basic api endpoints used in the application
package api

import (
	"redis_user_management/controller"

	"github.com/gin-gonic/gin"
)

// Function AppEndpoints adds and returns all the endpoints to the gin engine
func AppEndpints(engine *gin.Engine) *gin.Engine {

	user := engine.Group("/user")

	user.GET("/:id", controller.FetchUser)

	user.POST("/create", controller.CreateUser)

	user.PUT("/update", controller.UpdateUser)

	user.DELETE("/delete/:id", controller.DeleteUser)

	return engine
}
