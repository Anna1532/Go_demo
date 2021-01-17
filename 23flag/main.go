package main

import (
	"flag"
	"fmt"
	"time"
)

//flag获取命令行参数

func main(){
	//创建一个标志位参数,返回了一个指针
	name:=flag.String("name","Bob","Input name")
	age:=flag.Int("age",0,"输入年龄")
	married:=flag.Bool("married",false,"婚否")
	cTime:=flag.Duration("Time",time.Second,"多久了")//duration指的是时间间隔
	//使用flag都需要调用Parse
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
}
