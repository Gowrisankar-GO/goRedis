package api

import (
	"go_redis/controller"

	"github.com/gin-gonic/gin"
)

func AppEndpints(engine *gin.Engine) *gin.Engine {

	user := engine.Group("/user")

	user.GET("/:id", controller.FetchUser)

	user.POST("/create", controller.CreateUser)

	user.PUT("/update",controller.UpdateUser)

	user.DELETE("/delete/:id",controller.DeleteUser)

	return engine
}
