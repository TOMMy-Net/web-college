package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)


const (
	ServerError = "logging/server_errors.log"
	DbError     = "logging/db_errors.log"
)

func LoggerNew(f string) *logrus.Logger{
	var logrus = logrus.New()
	
	
	file, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.Out = file
	} else {
		logrus.Fatal("Failed to log to file, using default stderr")
		return nil
	}
	return logrus
	
}

