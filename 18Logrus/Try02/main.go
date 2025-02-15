package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 输出Json格式的日志，而不是ASCII格式
	log.SetFormatter(&log.JSONFormatter{})

	//输出到stdout而不是默认的stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// 只记录warning及以上级别的日志，Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}