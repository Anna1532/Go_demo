package main

import "fmt"

//输出
//func main(){
//	//path:="\"C:\\program\\Go\""//通过使用转义字符\来让系统识别“等要输出的符号
//	//fmt.Printf(path)
//	//fmt.Println(len(path))
//	name:="anna"
//	age:="22"
//	//fmt.Println(name+age)
//	s1:=fmt.Sprintf("%s%s",name,age)//Sprintf把拼接之后的字符串赋值给一个变量s1
//	fmt.Println(s1)
//	fmt.Printf("%s%s",name,age)//printf直接输出拼接的字符串
//}

//输入
func main(){
	var(
		name string
		age int
		married bool
	)
	//scan从标准输入扫描文本
	fmt.Scan(&name,&age,&married)//fmt.Scanf的用法与C语言类似
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}