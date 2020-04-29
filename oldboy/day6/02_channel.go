package main

import (
    "fmt"
)

/*
@Time : 2020/4/23 5:08 下午
@Author : audiRS7
@File : 02_channel
@Software: GoLand
*/
func main() {
    /*
       channel是引用类型，默认值为nil
    */

    // 错误演示，没有缓冲区，无法消费channel
    /*
    var ch1 chan  int
    ch1 <- 10
    ret := <- ch1
    fmt.Println(ret)*/
    // 正确演示
    ch2 := make(chan int,1) //chan 类型，接收类型，几个缓冲区
    ch2 <- 10
    ret := <- ch2
    fmt.Println(ret)
    close(ch2) //记得关闭
    /*
    1、关闭的通道可以再接收，但是值为0
    2、往关闭的通道中发送值，会发生panic
    3、关闭的通道再关闭，会发生panic
    */

}

