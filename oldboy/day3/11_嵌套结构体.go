package main

import (
    "fmt"
)
type student struct {
    name string
    age  int
    address //嵌套结构体
    email
}

type address struct {
    country string
    city    string
}
type email struct {
    city string
    e_address string
}

func main() {
    stu1 := student{
        name:"LiyaTong",
        age:28,
        address:address{
            country:"china",
            city:"BeiJing",
        },
        email:email{
            city:"Beijing",
            e_address:"jzb0424@163.com",
        },
    }

    fmt.Println(stu1.address.country)
    //fmt.Println(stu1.city)//错误实例：由于email和address结构体中均出现city，程序无法识别使用的是哪个结构体中的city
    //嵌套的结构体中如果有同名同类型的字段，需指定结构体的字段来使用
    fmt.Println(stu1.email.city)
}


