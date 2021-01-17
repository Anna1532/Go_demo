package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	//获取地址
 	resp,err:=http.Get("https://www.liwenzhou.com/")
 	if err != nil {
		fmt.Printf("Get failed!err:%v\n",err)
		return
	}
	//关闭回复的主体
	defer resp.Body.Close()
 	//读取body的内容
 	body,err:=ioutil.ReadAll(resp.Body)
 	if err != nil {
		fmt.Printf("Read from resp.Body failed!err:%v\n",err)
		return
	}
	fmt.Print(string(body))
 }
