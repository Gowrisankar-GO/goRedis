package controller

import (
	"log"
	"redis_user_management/logger"
)

var(
	InfoLog *log.Logger
	WarnLog *log.Logger
	ErrLog *log.Logger
)

func init(){

	InfoLog = loggger.InfoLog()

	WarnLog = loggger.WarnLog()

	ErrLog = loggger.ErrorLog()
}