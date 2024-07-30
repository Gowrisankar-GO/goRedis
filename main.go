// Package main is the entry point for the goRedis application.
package main

import (
	"log"
	"os"
	"redis_user_management/api"
	errPkg "redis_user_management/errors"

	_ "redis_user_management/docs" // To Import the generated docs

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func main() {

	appEngine := gin.Default()

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json")

	appEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initialize API endpoints
	Engine := api.AppEndpints(appEngine)

	port := os.Getenv("APP_PORT")

	if err := Engine.Run(":" + port); err != nil {

		log.Fatalf("%v", errPkg.ErrRunServer)
	}
}
