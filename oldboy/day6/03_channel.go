package main

import (
    "fmt"
)

/*
@Time : 2020/4/23 6:12 下午
@Author : audiRS7
@File : 03_channel
@Software: GoLand
*/
func recv(ch chan bool)  {
    ret := <- ch
    fmt.Println(ret)
}
func main() {
    ch := make(chan bool)
    go recv(ch)
    ch <- true
    fmt.Println("main 函数结束")
}


