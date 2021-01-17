package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func f1(){
	defer wg.Done()
	for i:=0;i<0;i++{
		fmt.Printf("A:%d\n",i)
	}
}
func f2(){
	defer wg.Done()
	for i:=0;i<10;i++{
		fmt.Printf("B:%d\n",i)
	}
}
func main(){
	runtime.GOMAXPROCS(2)//表示跑满6核心，默认是CPU的核心数
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
