package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, error := sql.Open("mysql", "root:1qaz@WSX3edc@rwebox.com@tcp(127.0.0.1:3306)/rwebox")
	if error != nil {
		fmt.Println("有错误")
	}
	fmt.Println(db)
}
