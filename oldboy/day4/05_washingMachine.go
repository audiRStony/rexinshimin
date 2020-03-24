package main

/*
@Time : 2020-03-24 10:34
@Author : audiRStony
@File : 05_washingMachine.go
@Software: GoLand
*/

import (
    "fmt"
)
//接口实现了洗衣机功能
//只要一个类型实现了wash(）和dry()方法，我们就称这个类型实现了xiyiji这个接口
type Xiyiji interface {
    Wash()
    Dry()
}

type Haier struct {
    name string
    price float32
    mode string
}

func (h Haier) Wash()  {
    fmt.Println("海尔洗衣机可以洗衣服")
}
func (h Haier) Dry()  {
    fmt.Println("海尔洗衣机可以甩干")
}

func main() {
    var a Xiyiji
    haier := Haier{
        name:"小神童",
        price:288,
        mode:"滚筒",
    }
    a = haier
    fmt.Println(a)
}

