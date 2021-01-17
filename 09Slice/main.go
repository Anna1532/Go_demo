package main

import "fmt"

func main(){
	//s1:=make([]int,5,10)//创建了一个int型长度为5.容量为10的切片
	s1:=[]int{0,1,2,3,4,5,6,7,8}
	//s2:=[]string{"Apple"}
	//删除元素的时候修改了原数组
	s1=append(s1[:3],s1[4:]...)//调用append函数必须要用原来的切片变量接受返回值
	//copy(s2,s1)
	fmt.Println(s1)
	//fmt.Printf("len:%d,cap:%d,%v",len(s1),cap(s1),s1)
	//var a1=[] int{1,3,5,7,9,11,13}
	//s1:=a1[0:4]//len()切片长度指的是切片的长度，而cap()切片容量指的是底层数组从切片的第一个元素开始的长度
	//s2:=s1[1:3]
	//fmt.Printf("len:%d,cap:%d,%d",len(s2),cap(s2),s2)
}


type statusType int

const (
	Norml statusType = 1<<iota
	Frezz
	
)
