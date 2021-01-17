package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(i int){
	fmt.Println("Hello",i)
	//开启一个单独的goroutine去执行hello()函数
	for i:=0;i<100;i++{
		//go hello(i) //哪个数字准备好了就打哪个，所以打印出来的数据的无序的
		go func(i int){
			fmt.Println(i)
		}(i)//不带参数时打印出来的数字有可能是重复的，需要去匿名函数外层获取i，启动goroutine需要一定的时间，而for循环会比较快，所以出现重复数字
	}
}
func f(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		r1:=rand.Int()
		r2:=rand.Intn(10)
		fmt.Println(r1,r2)
	}
}
func main(){
	i:=10
	hello(i)
}