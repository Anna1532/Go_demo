package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {//匿名函数
	defer func(){ //defer用来延迟运行，recover必须搭配defer使用
		err:=recover()//尝试恢复，尽量少用
		fmt.Println(err)
		fmt.Println("Doing...")
	}()
	panic("Error B")//出现panic之后会再运行一个defer，在defer中出现recover来恢复
	fmt.Println("still")//不执行
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	//funcA()
	funcB()
	//funcC()
}

