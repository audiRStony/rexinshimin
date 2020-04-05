package main

/*
@Time : 2020-04-04 10:57
@Author : audiRStony
@File : 03_反射ValueOf.go
@Software: GoLand
*/

import (
    "fmt"
    "reflect"
)

func reflectValueOf(x interface{}) {
    v := reflect.ValueOf(x) // 获取接口的值信息
    k := v.Kind()           // 拿到值对应的种类

    switch k {
    case reflect.Int64:
        fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
    // v.Int()从反射中获取整型的原始值，然后通过Int64()强制转化类型
    case reflect.Float32:
        fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
    // v.float32()从反射中获取float32的原始值，然后通过float32()强制转化类型
    case reflect.Float64:
        fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
        // v.float64()从反射中获取float64的原始值，然后通过float64()强制转化类型
    }
}

func main() {
    var a float32 = 13.14
    var b float64 = 100
    var c int64 = 1314
    reflectValueOf(a)
    reflectValueOf(b)
    reflectValueOf(c)
}
