package utility

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "log.txt",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     1, //days
	})
	return
}
