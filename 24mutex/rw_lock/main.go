package main

import (
	"fmt"
	"sync"
	"time"
)

//互斥锁可能导致性能下降
//读写互斥锁（读的操作远远大于写操作时使用）
var(
	x=0
	lock sync.Mutex
	rwlock sync.RWMutex
	wg sync.WaitGroup
)
func read(){
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	defer wg.Done()
}

func write(){
	rwlock.Lock()
	x=x+1
	time.Sleep(time.Millisecond*8)
	rwlock.Unlock()
	defer wg.Done()
}

func main(){
	start:=time.Now()
	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
