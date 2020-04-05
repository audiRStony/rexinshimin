package main

/*
@Time : 2020-03-09 23:00
@Author : audiRStony
@File : 04_结构体内存布局.go
@Software: GoLand
*/
import (
    "fmt"
)
//结构体中，每个属性的内存地址是连续的地址空间
func main() {
    type test struct {
        a int8
        b int8
        c int8
        d string
    }

    var tly = test{
        a : 9,
    }
    fmt.Println(&(tly.a))
    fmt.Println(&(tly.b))
    fmt.Println(&(tly.c))
    fmt.Println(&(tly.d))

}
