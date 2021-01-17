package main

import "fmt"

func main(){
	//var m1 map[string]int//其中string表示键的类型，int表示键对应的值的类型
	//fmt.Println(m1==nil)//未初始化，所以为空值，可以使用make来初始化
	m1:=make(map[int]int,10)//提前预估好map的容量，避免在程序运行期间扩容
	m1[3]=155
	//判断某个键是否存在
	v,ok:=m1[3]
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("no")
	}
	//fmt.Println(m1[3])
}
//遍历以及删除
func f1() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap,"张三")
	for k,v:= range scoreMap {
		fmt.Println(k,v)
	}
}
//元素类型为map的切片
func f2(){
	var s1=make([]map[int]string,5,10)
	s1[0]=make(map[int]string,1)
	s1[0][10]="apple"
	fmt.Println(s1[0][10])
	//值为切片类型的map
}