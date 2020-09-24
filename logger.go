package logger

import (
	// "fmt"
	"os"
	// "time"

	log "github.com/sirupsen/logrus"
)

// type StandardField struct {
// 	Name string
// 	Msg  string
// 	Time time.Time
// 	// "name": name,
// 	// 	"msg":  msg,
// 	// 	"time": time,
// }

// var (
// 	InfoLogger *log.Logger
// 	ErrorLogger *log.Logger
// )

var logging = log.New()

// type StandardField struct{
// 	Name string
// 	Msg string
// 	Time time.Time
// }

// type LogFields struct {
// 	s *StandardField[]string
// }
// var s  *StandardField
// var requestLogger = log.WithFields(log.Fields{"Name: ": s.Name, "Msg: ": s.Msg, "Time: ": s.Time})

// func Fields(name string, msg string, time time.Time) *log.Fields{

// 	log.SetFormatter(&log.JSONFormatter{})
// 	standardFields := log.Fields{
// 		// "Name": s.Name,
// 		// "Msg": s.Msg,
// 		// "Time": s.Time,
// 	}
// 	log.WithFields(standardFields)
// 	return &standardFields
// }



//Stand: 23.09.2020

//SetOutputFile ...
func SetOutputFile(loggingTo string){
	logFile, err := os.OpenFile(loggingTo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logging.Out = os.Stdout

	CheckErr(err)
	logging.Out = logFile
}

// // InfoLog makes new Info-Level Input in logfile
// func (s *StandardField) InfoLog(name string, msg string) {
	
// 	// log.Println(InfoLoggerName)
// }

// //ErrorLog makes new Error-Level input in logfile
// func ErrorLog(err error) {
// 	if err != nil {
// 		// log.NewEntry(ErrorLogger).Error()
// 	}
// }

// Formatter ...
func Formatter(a bool, colours bool) string{
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = time.Now().String()
	formatter.FullTimestamp = a
	formatter.ForceColors = colours

	log.SetFormatter(formatter)
	return formatter.TimestampFormat
}

func LevelInfo(msg string, name string, time string){
	log.Info("Job: " + name, " Msg: " + msg , " Time: " + time)
}

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
