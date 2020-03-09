package main

import (
	"fmt"
)

/*
@Time : 2020-03-09 17:50
@Author : audiRStony
@File : 01_homework.go
@Software: GoLand
*/

func main() {
	studentMap := make(map[string]map[string]int, 10)
	//初始化内层map
	studentMap["迪丽热吗"] = make(map[string]int, 3)
	studentMap["迪丽热吗"]["id"] = 1
	studentMap["迪丽热吗"]["age"] = 18
	studentMap["迪丽热吗"]["score"] = 100

	studentMap["娜扎"] = make(map[string]int, 3)
	studentMap["娜扎"]["id"] = 2
	studentMap["娜扎"]["age"] = 19
	studentMap["娜扎"]["score"] = 99


	//根据id查询信息
	for k, v := range studentMap {
		id,ok := v["id"]
		if ok {
		    if id == 1{
		        fmt.Println("name:",k)
		        for k2,v2 := range v{
		            fmt.Println(k2,v2)
                }
            }
        }else {
            fmt.Println("查无此人")
        }
	}

}
