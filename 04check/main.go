package main

import (
	"errors"
	"fmt"
	"log"
)

func check(x int) error{//检查x的值是否符合要求
	if x<=0{
		return errors.New("x<=0")//若小于0就报错
	}
	return nil
}

func main(){
	x:=-9
	if err:=check(x);err==nil{
		x++
		fmt.Println(x)
		log.Println(x)
	}else{
		log.Fatalln(err)
	}
}

