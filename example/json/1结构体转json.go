package main

import (
	"encoding/json"
	"fmt"
	//"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

/*
@Time : 2020/6/26 2:36 下午
@Author : audiRS7
@File : 1什么是json
@Software: GoLand
*/

/*
{
"name":"于谦",
"age":50,
"rmb":12345.67
"gender":"male",
"hobby":["抽烟","喝酒","烫头"],
"wife:{"name":"王钢蛋","gender":"female"},
"aunts":[
{"name":"王铁蛋","gender":"female"},
{"name":"王铜蛋","gender":"female"},
{"name":"王银蛋","gender":"female"}
]
}
*/

type Person struct {
	Name   string
	Age    int
	Rmb    float64
	Gender string
	Hobby  []string
}

func main() {
	YuQian()
}

//1、序列化，将结构体转化为json
func YuQian() {

	person := Person{"于谦", 50, 12345.67, "male", []string{"抽烟", "喝酒", "烫头"}}
	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(string(bytes))
}
