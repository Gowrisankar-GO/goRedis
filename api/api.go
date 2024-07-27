package api

import (
	"go_redis/controller"

	"github.com/gin-gonic/gin"
)

func AppEndpints(engine *gin.Engine) *gin.Engine {

	user := engine.Group("/user")

	user.GET("/:id", controller.FetchUser)

	user.POST("/create", controller.CreateUser)

	return engine
}
