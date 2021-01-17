package main
//将输入的命令简单输出
import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)
func main() {
	app := &cli.App{
		Name:  "arguments",//显示在帮助中
		Usage: "arguments example",
		//调用该命令行程序时实际执行的函数
		Action: func(c *cli.Context) error {
			//c.NArg返回命令行的数量
			for i:=0;i<c.NArg();i++{
				fmt.Printf("%d:%s\n",i+1,c.Args().Get(i))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}