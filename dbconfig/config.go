package dbconfig

import (
	"context"
	"fmt"
	errPkg "go_redis/errors"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// the initial database connection
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
		return &redis.Client{}, errPkg.ErrDbConn
	}

	fmt.Println("cmd",rdbCmd)

	return rdb, nil

}
