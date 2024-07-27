package models

import (
	"context"
	"go_redis/dbconfig"
	errPkg "go_redis/errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type BasicDefs struct{
	DbConn     *redis.Client
	Ctx        context.Context
	CtxCancel  context.CancelFunc
}

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Age       uint   `json:"age"`
	RoleId    uint   `json:"roleId"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    gin.H  `json:"data"`
}

var (
	Rdb     *redis.Client 
	UserKey string
)

func init() {

	if err := godotenv.Load(); err != nil {

		log.Fatalf("%v", errPkg.ErrLoadEnv)
	}

	var connErr error

	Rdb, connErr = dbconfig.DatabaseConnection(context.Background())

	if connErr != nil {

		log.Fatalf("%v", errPkg.ErrDbConn)
	}

	UserKey = "user:id"

	exists, err := Rdb.Exists(context.Background(), UserKey).Result()

    if err != nil {

        log.Fatalf("Could not check if counter exists: %v", err)
    }

    if exists == 0 {

        err = Rdb.Set(context.Background(), UserKey, 0, 0).Err()

        if err != nil {

            log.Fatalf("Could not set initial counter value: %v", err)
        }
    }

}
