package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)


const (
	ServerError = "logging/server_errors.log"
	DbError     = "logging/db_errors.log"
)

func LoggerNew() *logrus.Logger{
	var log = logrus.New()   
	log.Formatter = &logrus.TextFormatter{
		DisableColors: false,
	}
	return log
}

func WriteLog(filePath string) *os.File{
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	return file
}
