package main

import "fmt"

//方法

type dog struct{
	name string
}
type person struct{
	name string
	age int
}
//构造函数
func newDog(name string)dog{
	return dog{
		name:name,
	}
}
func newPerson(name string,age int)person{
	return person{
		name:name,
		age:age,//构造函数中不能赋值，需要在main函数中赋值
	}
}
//方法是作用域特定类型的函数
//接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang(){//需要限定哪些类型可以调用该方法，所以接收者需要写在方法名前面
	fmt.Printf("%s:汪汪汪\n", d.name)
}
//值接收者不能直接调用，方法中是外部的值，是copy过去的，调用之后值不改变，可以使用指针接收者
func (p *person) add(){
	p.age++
}
func main(){
	d1:=newDog("Anna")
	p1:=newPerson("Bob",18)
	d1.wang()//用.的方式调用方法
	fmt.Println(p1.age)
	p1.add()
	fmt.Println(p1.age)
}