package main

import(
	"fmt"
	"time"
)

func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func main(){
	//now:=time.Now()
	//fmt.Println(now.Month())
	timestampDemo()
}
