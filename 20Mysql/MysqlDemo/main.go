package main
import (
"database/sql"

_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库信息
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	//链接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
}