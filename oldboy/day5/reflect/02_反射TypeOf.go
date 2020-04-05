package main

/*
@Time : 2020-04-01 21:29
@Author : audiRStony
@File : 02_反射.go
@Software: GoLand
*/

import (
    "fmt"
    "reflect"
)

/*在Go语言的反射机制中，任何接口值都是由一个具体类型和具体类型的值两部分组成的
在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为
由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf
两个函数来获取任意对象的Type和Value*/

//类型反射
func reflectType(x interface{})  { //传递一个可以接收任何类型的空接口
    t := reflect.TypeOf(x)//拿到x的动态类型
    //fmt.Printf("type is : %v\n",t)
    fmt.Printf("name is %v,kind is %v\n",t.Name(),t.Kind())
    //fmt.Printf("%T\n",x)//本质就是反射
    /*ide的代码补全功能也是用的反射原理*/
}

type cat struct {
    name string
}

type person struct {
    name string
    age int
}

func main() {
    var c1 = cat{
        name:"XiaoHui",
    }
    var p1 = person{
        name:"Zzh",
        age:27,
    }
    //reflectType(100)
    reflectType(c1)
    reflectType(p1)

    var i int32 = 100
    var f float32 = 13.14
    /*对于指针类型(引用类型)，t.Name()为空*/
    reflectType(&i) //name is ,kind is ptr
    reflectType(&f) //name is ,kind is ptr
    reflectType(map[string]int{}) //name is ,kind is map
    reflectType([]int{1,2,3}) //name is ,kind is slice
}

