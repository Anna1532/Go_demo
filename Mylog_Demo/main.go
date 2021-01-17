package main

import (
	"Mylog_Demo/Mylog" //需要生成mod包，而且调用的包需要在一个运行目录之下
	"fmt"
	"math/rand"
	"time"
)

//测试自己完成的日志包

func main(){
	rand.Seed(time.Now().UnixNano())
	//"./"表示在此文件夹下构造一个新的日志文件，级别从Info开始，文件名称是20200707.log，文件大小是10*1024
	log:=Mylog.NewFileLog("debug","./",fmt.Sprintf("%v-%v",time.Now().Unix(),rand.Intn(1000))+".log",10*1024)//在文件中输出
	//在终端中构建日志信息，输出debug级别以下的日志
	//var log = Mylog.Newlog("debug")//在终端中输出
	for i:=0;i<10000;i++{
		//这里可以输出任何类型的信息，int，string...
		log.Debug("这是一个Debug日志")
		log.Trace("这是一个Trace日志")
		log.Info("这是一个Info日志")
		log.Warning("这是一个Warning日志")
		log.Error("这是一个Error日志")
		log.Fatal("这是一个Fatal日志")
		//time.Sleep(time.Second*2)
	}
}
