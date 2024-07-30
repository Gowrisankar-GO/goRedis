// Package main is the entry point for the goRedis application.
package main

import (
	"fmt"
	"log"
	"os"
	"redis_user_management/api"
	errPkg "redis_user_management/errors"

	_ "redis_user_management/docs" // To Import the generated docs

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	port := os.Getenv("APP_PORT")

	appEngine := gin.Default()

	docPath := fmt.Sprintf("http://localhost:%v/swagger/doc.json",port)

	url := ginSwagger.URL(docPath)

	appEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initialize API endpoints
	Engine := api.AppEndpints(appEngine)

	if err := Engine.Run(":" + port); err != nil {

		log.Fatalf("%v", errPkg.ErrRunServer)
	}
}
