package main
import "fmt"

func main(){
	i1 :=101
	fmt.Printf("%o\n",i1)//将二进制转换成八进制，%o表示八进制，%b表示二进制
	i2:=077
	fmt.Printf("%d\n",i2)//将八进制的数转换成了十进制，%d表示转换成十进制
	i3:=1001101
	fmt.Printf("%x\n",i3)//将二进制转换成十六进制，%x表示十六进制
}
