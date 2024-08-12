// Package errors contains basic errors where they are used in error handling in request handler functions
package info

import "errors"

var (
	ErrLoadEnv     = errors.New("failed to load env")
	ErrDbConn      = errors.New("failed to connect database")
	ErrRunServer   = errors.New("failed to run server")
	ErrSetUid      = errors.New("failed to initialize userid in redis database")
	ErrStructType  = errors.New("provided data is not a struct")
	ErrReqFld      = errors.New("required field is empty")
	ErrFldVal      = errors.New("invalid field value")
	ErrOpenLog     = errors.New("failed to open the log file")
	ErrLogFile     = errors.New("failed to create log file")
	ErrLogDir      = errors.New("failed to create log directory")
	ErrCreateUser  = errors.New("unable to create a new user")
	ErrDelUser     = errors.New("unable to delete the user")
	ErrUpdateUser  = errors.New("unable to update the user")
	ErrUkeyexist   = errors.New("could not check if userkey exists in db")
	ErrSetInitUkey = errors.New("could not set initial userkey value")
)