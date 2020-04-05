package main

/*
@Time : 2020-03-26 16:39
@Author : audiRStony
@File : 11_值接收和指针接收实现接口.go
@Software: GoLand
*/
import (
    "fmt"
)

type animal interface {
    speck()
    move()
}

type cat struct {
    name string
}

// 方法1：使用值接收者
/*func (c cat) speck()  {
    fmt.Println("miaomiaomiao")
}
func (c cat) move()  {
    fmt.Println("cat can move")
}*/

// 方法2：指针接收者
func (c *cat) speck() {
    fmt.Println("miaomiaomiao")
}
func (c *cat) move() {
    fmt.Println("cat can move")
}

func main() {
    var x animal

    hh := cat{"花花"} // cat
    x = hh /*实现animal接口的是指针类型*cat，此时hh是一个值类型cat，值类型是副本的拷贝*/

    tom := &cat{"汤姆猫"} // *cat
    tom.speck()        // (*tom).speck()
    x = tom
    fmt.Println(x)
}
