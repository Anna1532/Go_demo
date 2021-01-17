package main

import (
	"fmt"
	"reflect"
)
type D struct{}
func reflectType(x interface{}){
	//不知道别人会传进来什么类型的变量，所以参数类型定义为接口
	//1、通过类型断言
	//2、反射
	obj:=reflect.TypeOf(x)
	fmt.Println(obj,obj.Name(),obj.Kind())//Kind 是一种种类信息
	fmt.Printf("%T\n",obj)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main(){
	var a float32=1.234
	var b int=7
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
	//reflectType(a)
	//reflectType(b)
	//var c []int64
	//reflectType(c)
}