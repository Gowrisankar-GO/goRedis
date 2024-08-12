// Package dbconfig used to initialize a basic Redis database connection via env file
package dbconfig

import (
	"context"
	"fmt"
	"os"
	"redis_user_management/info"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// Function used to initialize redis database connection
func DatabaseConnection(ctx context.Context) (*redis.Client, error) {

	urlAddr := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := strconv.Atoi(os.Getenv("DB"))

	if err != nil {
		return &redis.Client{}, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     urlAddr,
		Password: os.Getenv("DB_PASSWORD"),
		DB:       db,
	})

	rdbCmd := rdb.Ping(ctx)

	if rdbCmd.Val() != "PONG" {
		return &redis.Client{}, info.ErrDbConn
	}

	fmt.Println("cmd", rdbCmd)

	return rdb, nil

}
