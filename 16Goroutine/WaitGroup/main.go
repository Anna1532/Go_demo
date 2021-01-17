package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
//使用了sync.WaitGroup来实现goroutine的同步，即等待所有的goroutine都结束之后再退出
var wg sync.WaitGroup

func f1(i int){
	defer wg.Done()
	time.Sleep(time.Microsecond *time.Duration(rand.Intn(200)))
	fmt.Println(i)
}

func main(){
	for i:=0;i<10;i++{
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()//等待wg的计数器减为0
}