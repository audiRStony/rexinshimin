package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
@Time : 2020/6/26 5:17 下午
@Author : audiRS7
@File : 9结构体切片写入json文件
@Software: GoLand
*/

type Person struct {
	Name   string
	Age    int
	Rmb    float64
	Gender string
	Hobby  []string
}

func main() {
	//encodejson2File()
	decodejson2Slce()
}
func encodejson2File() {
	p1 := Person{"王钢蛋", 30, 123.45, "female", []string{"抽中华", "喝可乐", "吃大蒜"}}
	p2 := Person{"王铜蛋", 30, 123.45, "female", []string{"抽中华", "喝雪碧", "吃大蒜"}}
	p3 := Person{"王铁蛋", 30, 123.45, "female", []string{"抽玉溪", "喝怡宝", "吃大葱"}}

	people := make([]Person, 0)
	people = append(people, p1, p2, p3)

	dstFile, _ := os.OpenFile("./badayi.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	defer dstFile.Close()
	//将go语言数据编码到json文件
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("enconding successful")
}

func decodejson2Slce() {
	people := make([]Person, 0)
	srcFile, _ := os.Open("./badayi.json")

	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(&people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
