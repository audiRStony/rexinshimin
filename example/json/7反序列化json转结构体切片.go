package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 4:51 下午
@Author : audiRS7
@File : 7反序列化json转结构体切片
@Software: GoLand
*/

func main() {
	type Person struct {
		Name   string
		Age    int
		Rmb    float64
		Gender string
		Hobby  []string
	}
	jsonStr := `[{"hobby":["抽中华","喝牛栏山","烫花卷头"],"name":"王钢蛋"},{"hobby":["抽玉溪","喝红星","烫杀马特"],"name":"王铁蛋"},{"hobby":["抽黄鹤楼","喝茅台","烫鸡冠头"],"name":"王银蛋"}]`
	jsonBytes := []byte(jsonStr)
	personSlice := make([]Person, 0)
	err := json.Unmarshal(jsonBytes, &personSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(personSlice)
}
