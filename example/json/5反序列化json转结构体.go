package main

import (
	"encoding/json"
	"fmt"
)

/*
@Time : 2020/6/26 3:28 下午
@Author : audiRS7
@File : 5反序列化json转结构体
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
	jsonStr := `{"Name":"于谦","Age":50,"Rmb":12345.67,"Gender":"male","Hobby":["抽烟","喝酒","烫头"]}`
	jsonBytes := []byte(jsonStr)
	personPrt := new(Person)
	err := json.Unmarshal(jsonBytes, personPrt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*personPrt)
}
