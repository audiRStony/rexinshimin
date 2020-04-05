package main

/*
@Time : 2020-04-04 16:50
@Author : audiRStony
@File : 04_reflectElem.go
@Software: GoLand
*/

import (
    "fmt"
    "reflect"
)

func modifyValue(x interface{})  {
    /*反射接收的是动态的类型以及类型的值
    原始类型为int64，所以修改时，只能修改为同类型不同值，不可跨类型
    */

    v := reflect.ValueOf(x)
    fmt.Println(v.Kind())

    if v.Kind() == reflect.Ptr { //判断kind是否为指针
        v.Elem().SetInt(12) //Elem()方法用来修改指针地址，提供了修改值的方法
    }
}
func main() {
    var a int64 = 100
    modifyValue(&a)
    fmt.Println(a)
}

