package main

/*
@Time : 2020-03-31 16:26
@Author : audiRStony
@File : 01_interface01.go
@Software: GoLand
*/
import (
    "fmt"
)
//接口是引用类型
//接口值 由两部分组成：类型和值
func main() {
    var x interface{} //<type,value>
    var a int64 = 10
    var b int32 = 20
    var c bool = true
    x = a // <int64,10>
    x = b //<int32,20>
    x = c //<bool,true>
    fmt.Println(x)
}

