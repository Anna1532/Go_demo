package main

import (
	"fmt"
	"io"
	"os"
)
func test1(){
	//只读的方式打开
	fileObj,err:=os.Open("D:/ProgramData/Go/src/hello/file/main.go")
	if err!=nil{
		fmt.Println("open file faild,err:%v",err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	//读文件
	for {
		var tmp=make([]byte,128)
		//切片，指定读取长度
		n,err:=fileObj.Read(tmp)
		if err==io.EOF{
			fmt.Println("读完了")
			return
		}
		if err!=nil{
			fmt.Println("Read file faild,err:%v",err)
			return
		}
		fmt.Printf("\n读了%d个字节\n",n)
		fmt.Println(string(tmp[:n]))
		if n<128{
			return
		}
	}
}
func main(){
	fileObj,err:=os.Open("D:/ProgramData/Go/src/hello/file/main.go")
	if err!=nil{
		fmt.Printf("Can't open this file,err:%v\n",err)
		return
	}
	//文件对象的类型
	fmt.Printf("%T\n",fileObj)
	//获取文件对象详细信息
	fileInfo,err:=fileObj.Stat()
	if err!=nil{
		fmt.Printf("Get fileInfo failed,err:%v\n",err)
		return
	}
	fmt.Printf("文件大小是:%dB\n",fileInfo.Size())
}