package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	apiurl:="http://127.0.0.1:9090/get"
	data:=url.Values{}
	data.Set("Name","Anna")
	data.Set("age","18")
	u,err:=url.ParseRequestURI(apiurl)//解析成url结构
	if err != nil {
		fmt.Printf("parse url requesturl failed,err:%v\n",err)
		return
	}
	u.RawQuery=data.Encode()//url编码
	fmt.Println(u.String())
	resp,err:=http.Get(u.String())
	if err != nil {
		fmt.Printf("Post failed,err:%v\n",err)
		return
	}
	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
}
