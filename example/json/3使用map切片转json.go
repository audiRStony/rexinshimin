package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 3:18 下午
@Author : audiRS7
@File : 3使用map切片转json
@Software: GoLand
*/

//3、使用map[string]interface{} 描述于谦小姨子并转json
func main() {

	dataMap1 := make(map[string]interface{})
	dataMap1["name"] = "王钢蛋"
	dataMap1["hobby"] = []string{"抽中华", "喝牛栏山", "烫花卷头"}

	dataMap2 := make(map[string]interface{})
	dataMap2["name"] = "王铁蛋"
	dataMap2["hobby"] = []string{"抽玉溪", "喝红星", "烫杀马特"}

	dataMap3 := make(map[string]interface{})
	dataMap3["name"] = "王银蛋"
	dataMap3["hobby"] = []string{"抽黄鹤楼", "喝茅台", "烫鸡冠头"}

	dataSlice := make([]map[string]interface{}, 0)
	dataSlice = append(dataSlice, dataMap1, dataMap2, dataMap3)

	bytes, err := json.Marshal(dataSlice)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(string(bytes))

}
