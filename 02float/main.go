package main

import "fmt"

func main(){
	//math.MaxFloat32//32位浮点数最大值
	f1:=1.3145456
	fmt.Printf("%T",f1)//go语言中小数默认都是float64的
	f2:=float32(1.3145456)//显示声明float是32位的
	fmt.Printf("\n%T",f2)//float32和float64是不同的类型不能相互赋值
}
