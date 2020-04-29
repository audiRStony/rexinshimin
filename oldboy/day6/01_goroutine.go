package main

import (
    "fmt"
    "runtime"
    "sync"
)

/*
@Time : 2020/4/23 3:03 下午
@Author : audiRS7
@File : 01_goroutine
@Software: GoLand
*/

var wg sync.WaitGroup
func hello()  {
    runtime.GOMAXPROCS(1) //若不写，默认是跑满cpu核心
    fmt.Println("hello 娜扎")
    wg.Done()// 没有done 会产生deadlock
}

func main() {
    defer fmt.Println("三元桥")
    wg.Add(1) // 视启动几个goroutine决定
    go hello()
    fmt.Println("hello main func")
    /*主函数执行完毕后退出，导致goroutine未创建或者未执行函数销毁未打印hello()函数内容*/

    /*不要随意在代码里使用sleep*/
    //time.Sleep(time.Second) //sleep 1秒让goroutine有cpu时间片来运行hello()
    wg.Wait()// 阻塞。等待goroutine执行完毕
}
