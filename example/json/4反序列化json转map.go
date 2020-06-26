package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 3:21 下午
@Author : audiRS7
@File : 4反序列化json转map
@Software: GoLand
*/

//将谦哥的json信息转为map
func main() {
	jsonStr := `{"Name":"于谦","Age":50,"Rmb":12345.67,"Gender":"male","Hobby":["抽烟","喝酒","烫头"]}`
	jsonBytes := []byte(jsonStr)
	dataMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &dataMap) //注意 第二个参数interface{}必须取地址
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	fmt.Println(dataMap)

}
