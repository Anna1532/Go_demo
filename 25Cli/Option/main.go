package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		//设置flag字段添加选项
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang", //设置想要输出的语言选项 Option.exe -lang English
				Value: "English",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "world" //name设置默认为world
			//如果输入了name，就输出用户输入的name Option.exe -lang English Anna
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "English" {
				fmt.Println("Hello", name)
			} else {
				fmt.Println("你好",name)
			}
			return nil
		},
	}
	err:=app.Run(os.Args)//调用run（）方法，传入命令行参数
	if err!=nil{
		log.Fatal(err)
	}
}
