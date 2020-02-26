package main

import (
    _"github.com/go-sql-driver/mysql"
    "fmt"
    "github.com/jmoiron/sqlx"
)

type Person struct {
    UserId int `db:"userid"`
    UserName string `db:"username"`
    Sex string `db:"sex"`
    Email string `db:"email"`
}

type Place struct {
    Counttry string `db: "country"`
    City string `db:"city"`
    TelCode int `db:"telcode"`
}

var Db *sqlx.DB


//初始化数据库登陆信息
func init()  {
    database,error := sqlx.Open("mysql","fuck:123456@tcp(localhost:3306)/Golang")
    if error != nil{
        fmt.Println("connect mysql failed")
        return
    }
    Db = database
}

func main()  {
    var person []Person

    err1 := Db.Select(&person,"select userid, username, sex, email from person where userid=?", 1)
    if err1 != nil{
        fmt.Println("select failed",err1)
        return
    }

    fmt.Println("select successfully\n",person)
}