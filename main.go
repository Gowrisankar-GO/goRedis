package main

import (
	"go_redis/api"
	errPkg "go_redis/errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	appEngine := gin.Default()

	Engine := api.AppEndpints(appEngine)

	port := os.Getenv("APP_PORT")

	if err := Engine.Run(":" + port); err != nil {

		log.Fatalf("%v", errPkg.ErrRunServer)
	}
}
