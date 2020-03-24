package main

/*
@Time : 2020-03-24 22:28
@Author : audiRStony
@File : 07_typejudge.go
@Software: GoLand
*/

import (
    "fmt"
)

func showtype07(x interface{}) {
    v1, ok := x.(int)
    if ok {
        fmt.Println("Int Type", v1)
    }

    v2, ok := x.(string)
    if ok {
        fmt.Println("String Type", v2)
    }
}

func Justiftype(a interface{}) {
    // 使用switch方式断言类型，可使用x.(type)来替代
    // 在case中写入需断言的类型即可
    //推荐使用switch方式
    switch v := a.(type) {
    case string:
        fmt.Printf("it's [string] type %v\n", v)
    case int:
        fmt.Printf("it's [int] type %v\n", v)
    case bool:
        fmt.Printf("it's [bool] type %v\n", v)
    default:
        fmt.Println("unsupport type!")
    }
}

func main() {
    // if 判断版本
    showtype07(100)
    showtype07("后羿")

    // switch判断版本
    Justiftype(20.2)
    Justiftype("黄忠")
}
