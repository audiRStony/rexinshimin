package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 3:17 下午
@Author : audiRS7
@File : 2使用map转json
@Software: GoLand
*/
//2、使用map[string]interface{}描述于谦并转json
func main() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 50
	dataMap["gender"] = "male"
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}
	bytes, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(string(bytes))
}
