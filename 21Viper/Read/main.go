package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){

	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	//设置默认值
	viper.SetDefault("name", "dogger")
	viper.SetDefault("age", "18")
	viper.SetDefault("class", map[string]string{"class01": "01", "class02": "02"})

	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。

	fmt.Println(
		viper.Get("name"), // 过去配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
	viper.WriteConfigAs("server-01.yaml")
}