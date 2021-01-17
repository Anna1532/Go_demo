package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	logrus.SetOutput(os.Stdout)
	// 设置日志等级
	logrus.SetLevel(logrus.InfoLevel)
}
func main() {
	logrus.WithFields(logrus.Fields{
		"test": "former",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
//显示级别以及消息
	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		//是一个接口，可以输出任何类型的内容
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}