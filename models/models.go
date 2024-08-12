// Package models contains init functions and the commonly required models to interact while reading and writing in database operations
package models

import (
	"context"
	"log"
	"redis_user_management/dbconfig"
	"redis_user_management/info"
	loggger "redis_user_management/logger"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type BasicDefs struct {
	DbConn    *redis.Client
	Ctx       context.Context
	CtxCancel context.CancelFunc
}

type User struct {
	Id        uint   `json:"id" validate:"omitempty"`
	FirstName string `json:"firstName" validate:"required,firstname"`
	LastName  string `json:"lastName" validate:"omitempty,lastname"`
	Email     string `json:"email" validate:"required,email"`
	Mobile    string `json:"mobile" validate:"required,mobile"`
	Age       uint   `json:"age" validate:"required,age"`
	RoleId    uint   `json:"roleId" validate:"required,roleid"`
	Status    uint   `json:"status" validate:"required,status"`
}

type Response struct {
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Data    map[string]interface{} `json:"data"`
}

var (
	Rdb           *redis.Client
	UserKey       string
	ValidatorKeys map[string]string
)

// Function init is a special type function used to execute database initialization and the initialization of basic application models needed for the smooth runninng of the application
func init() {

	if err := godotenv.Load(); err != nil {

		log.Fatalf("%v", info.ErrLoadEnv)
	}

	var connErr error

	Rdb, connErr = dbconfig.DatabaseConnection(context.Background())

	if connErr != nil {

		log.Fatalf("%v", info.ErrDbConn)
	}

	UserKey = "user:id"

	exists, err := Rdb.Exists(context.Background(), UserKey).Result()

	if err != nil {

		loggger.ErrorLog().Printf("%v: %v",info.ErrUkeyexist,err)

		log.Fatalf("%v: %v",info.ErrUkeyexist,err)
	}

	if exists == 0 {

		err = Rdb.Set(context.Background(), UserKey, 0, 0).Err()

		if err != nil {

			loggger.ErrorLog().Printf("%v: %v",info.ErrSetInitUkey,err)

			log.Fatalf("%v: %v", info.ErrSetInitUkey, err)
		}
	}

	ValidatorKeys = map[string]string{
		"Tag":       "validate",
		"Required":  "required",
		"Omit":      "omitempty",
		"FirstName": "firstname",
		"LastName":  "lastname",
		"Email":     "email",
		"Mobile":    "mobile",
		"Age":       "age",
		"Role":      "roleid",
		"Status":    "status",
	}

}

// Function ConvertStructToMap used to return a struct typed data into a map[string]interface{} data
func ConvertStructToMap(data interface{}) (map[string]interface{}, error) {

	v := reflect.ValueOf(data)

	// confirming whether the provided data is a struct type or not

	if v.Kind() != reflect.Struct {

		return map[string]interface{}{}, info.ErrStructType
	}

	t := reflect.TypeOf(data)

	mapData := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {

		fieldValue := v.Field(i)

		tagName := t.Field(i).Tag.Get("json")

		mapData[tagName] = fieldValue.Interface()

	}

	return mapData, nil

}
