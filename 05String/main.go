package main
import "fmt"
func main(){
	s:="hello12334"
	//fmt.Println(len(s))
	//for i:=0;i<len(s);i++{
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n",s[i]
	//}
	//for _,c:=range s{
	//	fmt.Printf("%c\n",c)
	//}
	s1:=[]rune(s)//将字符串强制转换成一个rune切片
	s1[0]='H'//将字符串s中的第一个字母转换成大写
	fmt.Printf("%c\n",s1[0])
	fmt.Println(string(s1))
}
