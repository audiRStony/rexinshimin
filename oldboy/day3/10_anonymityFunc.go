package main

/*
@Time : 2020-03-14 15:14
@Author : audiRStony
@File : 10_anonymityFunc.go
@Software: GoLand
*/

import (
    "fmt"
)

type student struct {
    name string
    //匿名字段。类型即为字段名
    string
    int
}

func main() {
    stu1 := student{
        name : "LiyaTong",

    }
    fmt.Println(stu1.name)
    fmt.Println(stu1.string)
    fmt.Println(stu1.int)

}

