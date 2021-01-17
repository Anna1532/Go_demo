package main

import "fmt"

//函数定义
//没有返回值
/*func f1(x int,y int){
	fmt.Println(x+y)
}
//没有参数
func f2()int{
	ret:=3
	return ret
}
//参数+返回值
func f3(x int,y int)(ret int){
	ret=x+y
	return
}
//可变长参数,...表示可以有多个参数，也可以没有参数，可变长参数必须放在最后
func f4(x string,y...int){
	fmt.Println(x)
	fmt.Println(y)//y的类型是切片
}
//闭包
func f5(f func()){
	fmt.Println("This is f5")
	f()
}
func f6(x,y int){
	fmt.Println("This is f6")
	fmt.Println(x+y)
}
func main(){
	err:=errors.New("can't read")

}*/

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}