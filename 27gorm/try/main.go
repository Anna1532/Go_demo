package main

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

}