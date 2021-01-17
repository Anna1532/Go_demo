package main

import "fmt"

func main(){
	//i:=50
	//if i<18{
	//	fmt.Println("Kids")
	//}else if i>=18&&i<=40{
	//	fmt.Println("Adults")
	//}else {
	//	fmt.Println("Old")
	//}
	s:="hello12345"
	fmt.Println(len(s))
	for _,v:=range s{//i表示索引值，v表示值
		fmt.Printf("%c\n",v)
	}
}
