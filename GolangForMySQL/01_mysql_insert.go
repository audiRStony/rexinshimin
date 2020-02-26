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
    //插入信息
    input,err1 :=Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "female", "stu02@qq.com")
    if err1 != nil{
        fmt.Printf("Insert data failed %v\n",err1)
        return
    }

    //插入成功后返回id，以此判断是否插入成功
    id,err2 := input.LastInsertId()
    if err2 != nil{
        fmt.Println("exec failed",err2)
        return
    }

    fmt.Println("insert success",id)
}
