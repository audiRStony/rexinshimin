package main

/*
@Time : 2020-03-14 16:17
@Author : audiRStony
@File : 12_json.go
@Software: GoLand
*/

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    //默认json解析的首字母为大写，如需小写，在字段后使用json tag即可。
    //意为json包来处理的时候，按json tag的要求执行
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Gender string `json:"gender"`
}

func main() {
    stu1 := Person{
        ID:     1,
        Name:   "LiyaTong",
        Gender: "female",
    }
    /*序列化*/
    v, err := json.Marshal(stu1) // 返回的v为[]byte
    if err != nil {
        fmt.Printf("json格式化失败 %s\n", err)
    }

    fmt.Println(string(v)) // 使用string做强制转化

    /*反序列化*/
    str := `{"ID":1,"Name":"LiyaTong","Gender":"female"}`
    stu2 := &Person{}
    json.Unmarshal([]byte(str), stu2)
    fmt.Println(stu2)
    fmt.Printf("%T\n", stu2)
}
