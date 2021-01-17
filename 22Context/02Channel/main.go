package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var exitChan chan bool=make(chan bool,1)

func f(){
	defer wg.Done()
	LOOP:
	for{
		fmt.Println("Anna")
		time.Sleep(time.Millisecond*500)
		select{
		case <-exitChan:
			break LOOP//跳出for循环
		default:
		}
	}
}

func main(){
	//如何通知子goroutime退出
	wg.Add(1)
	go f()
	time.Sleep(time.Second*3)//3s之后通知goroutine退出
	exitChan<-true
	wg.Wait()
}
