package logger

import (
	// "fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func Fields(name string, msg string, time time.Time) {

	log.SetFormatter(&log.JSONFormatter{})
	standardFields := log.Fields{
		"name": name,
		"msg":  msg,
		"time": time,
	}
	log.WithFields(standardFields)
}

var logging = log.New()

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func SetOutputFile(loggingTo string) string {
	logFile, err := os.OpenFile(loggingTo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logging.Out = os.Stdout

	CheckErr(err)
	logging.Out = logFile
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	return logFile.Name()
}

func ErrorLog(err error) {
	if err != nil {
		log.Error(err)
	}
}

func CheckErr(err error) {
	if err != nil {

	}
}
