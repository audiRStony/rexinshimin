package main

/*
@Time : 2020-03-13 18:02
@Author : audiRStony
@File : 09_method.go
@Software: GoLand
*/

/*
可以给任意类型追加方法
不能给别的包定义的类型添加方法
只能给本包里的类型添加方法
*/

import (
    "fmt"
)

type Myint int

/*错误示范
//此函数会导致原int类型无法使用
func (i *int)sayHello()  {

}
*/


func (m *Myint) sayHi()  {
    fmt.Println("hello Myint")
}

func main() {
    var a Myint
    fmt.Println(a)
    a.sayHi()
}