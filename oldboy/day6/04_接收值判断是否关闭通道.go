package main

import (
	"fmt"
)

/*
@Time : 2020/5/12 9:59 下午
@Author : audiRS7
@File : 04_接收值判断是否关闭通道
@Software: GoLand
*/

func send(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch) //程序运行时，管道非必须显示关闭，可被gc垃圾回收机制回收
}
func main() {
	ch1 := make(chan int, 100)
	go send(ch1)
	//利用for循环去通道ch1取值
	/*    for  {
	      ret,ok := <- ch1 //使用value,ok := <- ch1 取值方式，当通道关闭是，ok为false
	      if !ok{
	          break
	      }
	      fmt.Println(ret)
	  }*/

	//利用for range取值
	for ret := range ch1 { //同ret,ok方式，range取不到值时，自动ok为false
		fmt.Println(ret)
	}
}
