package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main(){
	add:=&cli.App{
		Commands:[]*cli.Command{
			{
				Name: "add",//命令参数
				Aliases:[]string{"a"},
				Usage: "add a task to a list",
				Action:func(c *cli.Context)error{
					fmt.Println("added task:",c.Args().First())
					return nil
				},
			},
			{
				Name: "complete",
				Aliases:[]string{"c"},
				Usage: "complete a task to a list",
				Action:func(c *cli.Context)error{
					fmt.Println("complete task:",c.Args().First())
					return nil
				},
			},
			{
				Name: "template",
				Aliases:[]string{"c"},//命令参数
				Usage: "options for task templates",
				Subcommands:[]*cli.Command{
					{
						Name: "add",
						Usage: "add a new template",
						Action:func(c *cli.Context)error{
							fmt.Println("new task template:",c.Args().First())
							return nil
						},
					},
					{
						Name: "remove",
						Usage: "remove an existing template",
						Action:func(c *cli.Context)error {
							fmt.Println("remove task template:", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}
	err:=add.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}
