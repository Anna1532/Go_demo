package main
import (
	"encoding/json"
	"fmt"
)
type person struct{//json包中需要用到这些数据，所以需要以大写祖母开头
	Name string  `json:"name" db:"name" ini:"name"`	//如果需要在前端等地方小写。可以后面规定
	Age int		`json:"age"`//下次学反射
}

func main(){
	p1:=person{
		Name:"Anna",
		Age:19,
	}
	//序列化
	b,err:=json.Marshal(p1)	//Go语言中的结构体变量---->json格式的字符串
	if err!=nil{
		fmt.Printf("marshal faild,err:%v",err)
		return
	}
	fmt.Println(p1)
	fmt.Printf("%v\n",string(b))//如果不报错就输出json格式的结构体内容
	//反序列化
	str:=`{"name":"Anna","age":18}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)//指针是为了把值传递到json中
	fmt.Printf("%#v\n",p2)
}
