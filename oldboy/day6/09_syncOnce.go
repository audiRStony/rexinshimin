package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2020/5/24 2:55 下午
@Author : audiRS7
@File : 09_syncOnce
@Software: GoLand
*/

/*
sync.Once是golang package 中使方法只执行一次的对象实现，作用与init函数类似。但也有区别
1、init函数是在文件包首次被加载的时候执行，且只执行一次
2、sync.Once是在代码运行中需要的时候执行，且只执行一次
当一个函数不希望程序一开始的时候就被执行的时候，可以使用sync.Once
*/

var onlyOnce sync.Once

func f1(a int) {
	fmt.Println(a)
}
func closer(x int) func() {
	return func() {
		f1(x)
	}
}

func main() {
	f := closer(10) //使用闭包的方式，给函数传参且带入sync.Once
	onlyOnce.Do(f)  //只有一个Do方法，参数必须为函数且函数不能传参.其内部实现为一个互斥锁和一个布尔值，执行成功返回true，再有程序调用此数据时无需重复加载

}
