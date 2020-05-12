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

//无缓冲
/*func main() {
    ch := make(chan bool)
    go recv(ch)
    ch <- true
    fmt.Println("main 函数结束")
}*/

//有缓冲
func main()  {
    ch := make(chan bool,1)
    fmt.Println(len(ch),cap(ch)) //len获取长度，cap获取容量
    ch <- false
    go recv(ch)
    ch <- true
    fmt.Println("main函数运行结束")
}

