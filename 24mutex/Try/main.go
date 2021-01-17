package main

import (
	"fmt"
	"sync"
)
//互斥锁
var x=0
var wg sync.WaitGroup
var lock sync.Mutex
//不加锁的时候因为不同goroutine执行时机不同，所以每次得到的值不同
//对公共资源加锁后每次得到的值就相同了
func add(){
	defer wg.Done()
	for i:=0;i<100000;i++{
		lock.Lock()
		x++
		lock.Unlock()
	}

}
func main(){
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}