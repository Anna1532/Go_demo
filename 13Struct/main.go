package main

import "fmt"

//type people struct{//func后面跟着函数，type后面跟着类型，type用来新建一个类型
//	name,address string
//	age int
//	married bool
//	//hobby []string
//}
//func main(){
//	people.name="Anna"
//	a:=people{//创建结构体指针不需要单独创建结构体，只需要在main函数中加上&即可
//		name:"Anna",
//		address:"road",
//		age:18,
//		married:false,
//	}
//	//var a=new(people)
//
//	fmt.Println(a)
//	fmt.Printf("%T",a)
//}

//构造函数,通过调用构造函数来给结构体赋值
type person struct{
	name string
	age int
}
//构造函数约定俗成以new开头
func newPerson(name string,age int) *person{//当结构体中的内容较多时，可以把person换成指针
	return &person{
		name:name,
		age:age,
	}
}
func main(){
	p1:=newPerson("Anna",12)
	p2:=newPerson("bbb",19)
	fmt.Println(p1,p2)
}
