package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
@Time : 2020/6/26 4:58 下午
@Author : audiRS7
@File : 8读写json文件
@Software: GoLand
*/

func main() {
	//writejson2File()
	readFile2map()
}

func writejson2File() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 50
	dataMap["gender"] = "male"
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}
	//创建并打开目标json文件
	dstFile, _ := os.OpenFile("./yuqian.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	defer dstFile.Close()
	//将go语言数据编码到json文件
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("enconding successful")
}

func readFile2map() {
	srcFile, _ := os.Open("./yuqian.json")
	defer srcFile.Close()

	mapData := make(map[string]interface{})

	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(&mapData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decoding successfully")
	fmt.Println(mapData)
}
