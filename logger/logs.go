package loggger

import (
	"fmt"
	"log"
	"os"
	"redis_user_management/info"
)

const (
    InfoLevel = iota
    WarningLevel
    ErrorLevel
)

type Logger struct{
	Level    int
    Logger   *log.Logger
}

var(
	InfoLogger   Logger
	WarnLogger   Logger
	Errlogger    Logger
)

func init(){

	var LogFilePath =  "log/logger.log"

	_,err := os.Stat(LogFilePath)

	var Logfile *os.File

	if os.IsNotExist(err){

		err := os.MkdirAll("log",os.ModePerm)

		if err!=nil{

			fmt.Printf("%v",info.ErrLogFile)
		}

		Logfile,err = os.Create(LogFilePath)

		if err != nil{

			fmt.Printf("%v",info.ErrLogDir)
		}

		fmt.Println("successfully created the log file")
	}

	Logfile ,err = os.OpenFile(LogFilePath,os.O_APPEND|os.O_WRONLY,0577)

	if err != nil{

		fmt.Printf("%v",info.ErrOpenLog)
	}

	InfoLogger = Logger{Level: InfoLevel,Logger: log.New(Logfile,"INFO ",log.LstdFlags|log.Lshortfile)}

	WarnLogger = Logger{Level: WarningLevel,Logger: log.New(Logfile,"WARN ",log.LstdFlags|log.Lshortfile)}

	Errlogger = Logger{Level: ErrorLevel,Logger: log.New(Logfile,"ERROR ",log.LstdFlags|log.Lshortfile)}
}

func InfoLog() *log.Logger{

	return InfoLogger.Logger
}

func WarnLog() *log.Logger{
	
	return WarnLogger.Logger
}

func ErrorLog() *log.Logger{
	
	return Errlogger.Logger
}