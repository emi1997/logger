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

// type FieldHook struct {
// 	Name string
// 	Msg  string
// 	Time time.Time
// }

var logger = log.New()

func SetOutputFile(loggingTo string) {
	logFile, err := os.OpenFile(loggingTo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger.Out = os.Stdout

	checkErr(err)
	logger.Out = logFile
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// func NewFieldHook(name string, msg string, env string) *FieldHook {
// 	return &FieldHook{
// 		Name: name,
// 		Msg:  msg,
// 		Time: time.Time{},
// 	}
// }

// func (h *FieldHook)Levels() []logrus.Level{
// 	return logrus.AllLevels

// }

// func (h *FieldHook) Fire(entry *logrus.Entry)error {
// 	h.Msg = "Logged successfully!"

// 	entry.Data["Name"] = h.Name
// 	entry.Data["Msg"] = h.Msg
// 	entry.Data["Time"] = h.Time
// 	return nil
// }

// func main() {

// 	log.AddHook(NewFieldHook(Name, msg, config.Env))
// }

// func ErrorLog(loggingTo string) {
// 	SetOutputFile(loggingTo)
// 	// log.Fatalf()
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
