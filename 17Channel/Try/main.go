package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var a chan int
var b chan int

//无缓冲区通道
func NoBuffChannel(){
	//通过make来初始化通道
	a=make(chan int)
	wg.Add(1)
	go func(){
		defer wg.Done()
		//把通道a中的值赋值给x
		x:=<-a
		fmt.Println("从后台goroutine中取到了",x)
	}()//匿名函数一直在等待，当a<-1放进通道中时就立刻把值赋给a
	//把1放到通道a中
	a<-1
	fmt.Println("1已发送到通道a中")
	wg.Wait()
	close(a)
}

//带缓冲区的通道
func BuffChannel(){
	fmt.Println(b)//nil
	b=make(chan int,1)
	b<-35
	//这里b是一个通道地址
	fmt.Println("35已发送到通道中,b=",b)
	x:=<-b
	//把值从通道中取出来才能显示通道中的值
	fmt.Println("从通道b中取到了x",x)
	close(b)
}
func main(){
	BuffChannel()
}


