package main
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
//利用bufio读文件
func ReadbyBufio(){
	//只读的方式打开
	fileObj, err := os.Open("D:/ProgramData/Go/src/hello/file/main.go")
	if err != nil {
		fmt.Println("open file faild,err:%v", err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	reader:=bufio.NewReader(fileObj)
	for{
		line,err:=reader.ReadString('\n')//一行一行读
		if err==io.EOF{
			fmt.Println("读完了")
			return
		}
		if err!=nil{
			fmt.Println("Read file faild,err:%v",err)
			return
		}
		fmt.Print(line)
	}
}
func readbyIoutil(){
	ret,err:=ioutil.ReadFile("D:/ProgramData/Go/src/hello/file/main.go")
	if err != nil {
		fmt.Println("read file faild,err:%v", err)
		return
	}
	fmt.Print(string(ret))
}
//运行一次在文件中追加一次
func writeFile(){
	//以只写的方式打开，如果没有就创建一个，打开之后再追加
	fileObj,err:=os.OpenFile("D:/ProgramData/Go/src/hello/file/hey.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC,0644)
	if err!=nil{
		fmt.Println("open file faild,err:%v",err)
		return
	}
	fileObj.Write([]byte("Anna"))//写入字节切片数据
	fileObj.WriteString(" You are the best!")//直接写入数据
	fileObj.Close()
}
func main() {
	writeFile()
	readbyIoutil()
}