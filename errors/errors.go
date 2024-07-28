package errors

import "errors"

var (
	ErrLoadEnv = errors.New("failed to load env")
	ErrDbConn = errors.New("failed to connect database")
	ErrRunServer = errors.New("failed to run server")
	ErrSetUid  = errors.New("failed to initialize userid in redis database")
	ErrStructType = errors.New("provided data is not a struct")
)