package main

import "fmt"

func main(){
	ch:=make(chan int,1)
	for i:=0;i<10;i++{
		select{
		//第一次for循环：ch中没有值，所以不能传值给x，开始执行ch<-i，把i值传到ch中
		//第二次for循环：ch中有值了，将ch中的值传到x中输出，这时ch通道已满，i不能传值到ch中
		case x:=<-ch:
			fmt.Println(x)
		case ch<-i:

		}//0，2，4，6，8
	}
}
