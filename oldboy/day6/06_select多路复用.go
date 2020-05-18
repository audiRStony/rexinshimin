package main

import (
	"fmt"
	"math"
	"time"
)

/*
@Time : 2020/5/18 9:08 下午
@Author : audiRS7
@File : 06_select多路复用
@Software: GoLand
*/

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i <= math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f1:%d", i)
		time.Sleep(time.Millisecond * 50)
	}
}

func f2(ch chan string) {
	for i := 0; i <= math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f2:%d", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go f1(ch1)
	go f2(ch2)
	//var ch = make(chan int, 1)
	for {
		select {
		case ret := <-ch1:
			fmt.Println(ret)
		case ret := <-ch2:
			fmt.Println(ret)
		default:
			fmt.Println("暂时取不到值")
			time.Sleep(time.Millisecond * 1000)
		}
	}
}
