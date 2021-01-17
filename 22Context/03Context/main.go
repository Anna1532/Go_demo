package main
import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func f2(ctx context.Context){
	defer wg.Done()
LOOP:
	for{
		fmt.Println("Bob")
		time.Sleep(time.Millisecond*500)
		select{
		case <-ctx.Done()://调用ctx.Done(),只会返回一个只读的channel
			break LOOP//跳出for循环
		default:
		}
	}
}

func f(ctx context.Context){
	defer wg.Done()
	go f2(ctx)
LOOP:
	for{
		fmt.Println("Anna")
		time.Sleep(time.Millisecond*500)
		select{
		case <-ctx.Done()://调用ctx.Done(),只会返回一个只读的channel
			break LOOP//跳出for循环
		default:
		}
	}
}

func main(){
	//如何通知子goroutime退出
	ctx,cancel:=context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*3)//3s之后通知goroutine退出
	cancel()//调用cancel函数退出
	wg.Wait()
}
