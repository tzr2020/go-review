/*
	连接数据库
*/

package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Printf("连接数据库失败: %v\n", err)
	}

	log.Println("连接数据库成功")
}
