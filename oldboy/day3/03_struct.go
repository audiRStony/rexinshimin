package main

/*
@Time : 2020-03-09 22:04
@Author : audiRStony
@File : 03_struct.go
@Software: GoLand
*/
import (
    "fmt"
)

type Student struct {
    name string
    age int
    gender string
    hobby []string
}

func struct1()  {
    //实例化1
    //struct{}是值类型
    //如果初始化时没有给属性设定对应的值，那么对应的属性，其值为类型的默认值
    var syf = Student{}
    fmt.Println(syf.name)
    fmt.Println(syf.age)
    fmt.Println(syf.gender)
    fmt.Println(syf.hobby)
    /*
       string 默认值为空串
       int默认值为0
       []string默认值为[]
    */
}

func struct2()  {
    var LiyaTong = new(Student) //此处LiyaTong为内存地址。等同于var LiyaTong = &Student{}
    //(*LiyaTong).hobby
    /*在go语言中，指针是无法被编辑的，编辑的是指针对应的值,所以无需(*LiyaTong).hobby这种方式编辑，直接编辑即可 */
    LiyaTong.name = "佟丽娅"
    fmt.Println(LiyaTong.name)
}

func main()  {
    //struct1()
    struct2()
}

