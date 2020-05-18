package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2020/5/18 9:35 下午
@Author : audiRS7
@File : 07_ 并发控制
@Software: GoLand
*/
var x int64
var wg sync.WaitGroup
var lock sync.Mutex //定义一个互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //把厕所门关上
		/*1、从内存中找到x的值
		2、执行+1操作
		3、把（2）的结果再赋值给x写到内存*/
		x = x + 1
		lock.Unlock() //把厕所门打开
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
