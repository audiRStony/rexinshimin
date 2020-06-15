package mysql

/*
@Time : 2020/6/13 1:53 下午
@Author : audiRS7
@File : conn.go
@Software: GoLand
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ := sql.Open("mysql", "root:123456@tcp(192.168.106.130:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000) //最大连接数
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql.err:" + err.Error())
		os.Exit(1)
	}
}

//DBConn 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
