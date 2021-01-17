package main

import (
	"fmt"
	"sync"
)

//1、启动一个goroutine，生成100个数发送到ch1中
//2、启动一个goroutine，从ch1中取值，计算平方放到ch2中
//3、在main函数中，从ch2取值，打印出来

var wg sync.WaitGroup
//Produce函数只用来接收值，所以可以设置成单向通道
func Produce(ch1 chan<- int){
	defer wg.Done()
	for i:=0;i<100;i++{
		ch1<-i
	}
	close(ch1)
}

func Squre(ch1 <-chan int,ch2 chan<- int){
	defer wg.Done()
	for {
		x,ok:=<-ch1
		if !ok{
			break
		}
		ch2<-x*x
	}
	close(ch2)
}

func main(){
	//管道初始化
	a:=make(chan int,100)//一边放一边取出
	b:=make(chan int,100)
	wg.Add(2)//有n个goroutine，则wg.Add(n)
	go Produce(a)
	go Squre(a,b)
	wg.Wait()
	for ret:=range b{
		fmt.Println(ret)
	}
}