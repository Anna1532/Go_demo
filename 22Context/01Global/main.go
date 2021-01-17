package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var notify bool

func f(){
	defer wg.Done()
	for{
		fmt.Println("Anna")
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}

func main(){
	//如何通知子goroutime退出
	wg.Add(1)
	go f()
	time.Sleep(time.Second*3)//3s之后通知goroutine退出
	notify=true
	wg.Wait()
}
