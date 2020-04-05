package main

/*
@Time : 2020-04-04 17:26
@Author : audiRStony
@File : 05_结构体反射.go
@Software: GoLand
*/
import (
    "fmt"
    "reflect"
)

type student struct {
    Name string `json:"name"`
    Score int `json:"score"`
}

func main() {
    stu1 := student{
        Name:"娜扎",
        Score:88,
    }
    t := reflect.TypeOf(stu1)
    fmt.Println(t.Name(),t.Kind()) //student struct
    //通过for循环遍历结构体所有字段
    for i := 0; i < t.NumField();i ++ {
        field := t.Field(i)
        fmt.Printf("Name : %s ,Index : %d ,type : %v , json tag : %v\n",field.Name,field.Index,field.Type,field.Tag.Get("json"))
    }

    /*通过字段名获取指定结构体字段的信息*/
    if scoreField,ok := t.FieldByName("Score");ok {
        fmt.Printf("Name : %s Index : %d Type : %v json tag : %v\n",scoreField.Name,scoreField.Index,scoreField.Type,scoreField.Tag.Get("json"))
    }
}

