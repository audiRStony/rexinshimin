package main

/*
@Time : 2020-03-09 23:17
@Author : audiRStony
@File : 05_构造函数.go
@Software: GoLand
*/
import (
    "fmt"
)
type student struct {
    name string
    age  int
    gender string
    hobby []string
}

func main() {
    hobby := []string{"男人","女人"}
    newStudent("LiyaTong",18,"female",hobby)
    //newStudent2("LiyaTong",18,"female",hobby)

}

func newStudent(n string,a int,g string,h []string)   {
    s := student{
        name:n,
        age:a,
        gender:g,
        hobby:h,
    }
    fmt.Println(s)
}
func newStudent2(n string,a int,g string,h []string) *student  {
    return &student{
        name:n,
        age:a,
        gender:g,
        hobby:h,
    }
    //优点在于返回函数中返回指针，直接饮用结构体的地址，无需copy
}
