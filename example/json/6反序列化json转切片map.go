package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 4:34 下午
@Author : audiRS7
@File : 6json转切片
@Software: GoLand
*/

func main() {
	jsonStr := `[{"hobby":["抽中华","喝牛栏山","烫花卷头"],"name":"王钢蛋"},{"hobby":["抽玉溪","喝红星","烫杀马特"],"name":"王铁蛋"},{"hobby":["抽黄鹤楼","喝茅台","烫鸡冠头"],"name":"王银蛋"}]`
	jsonByte := []byte(jsonStr)
	jsonSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal(jsonByte, &jsonSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonSlice)
}
