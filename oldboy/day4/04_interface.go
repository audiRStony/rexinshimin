package main

/*
@Time : 2020-03-21 10:40
@Author : audiRStony
@File : 04_interface.go
@Software: GoLand
*/

import (
    "fmt"
)

type Cat struct{}

func (c Cat) Say() string {
    return "喵喵喵"
}

type Dog struct{}

func (d Dog) Say() string {
    return "汪汪汪"
}

type animal interface {
    Say() string
}

/*
type 接口类型名字 interface{
    方法名1(参数列表1) 返回列表1
    方法名2(参数列表2) 返回列表2
}
*/
// 接口名：使用type将接口定义为自定义的类型名，Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，
// 有字符串功能的接口叫Stringer等，接口名最好要能突出该接口的类型含义。
// 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在包(package)之外的代码访问
// 参数列表：返回值列表：参数列表和返回值列表中的参数变量名可以省略

func main() {
    // c := Cat{}
    // fmt.Println(c.Say())
    //
    // d := Dog{}
    // fmt.Println(d.Say())

    var animalList []animal
    c := Cat{}
    d := Dog{}
    animalList = append(animalList,c,d)
    //fmt.Println(animalList)
    for _, Items := range animalList{
        ret := Items.Say()
        fmt.Println(ret)
    }
}
