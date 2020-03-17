package main

/*
@Time : 2020-03-18 00:37
@Author : audiRStony
@File : 03_time_exercises.go
@Software: GoLand
*/
import (
    "fmt"
    "time"
    "strconv"
)

// 格式化打印
func timePrint(t time.Time) {
    timer := t.Format("2006/01/02 15:04:05")
    fmt.Println(timer)
}

// 测试用例
func exercises() {
    start_time := time.Now().Unix()
    str := "" // 声明一个字符串
    for i := 0; i < 100000; i++ { // for循环10W次拼接
        str += "golang" + strconv.Itoa(i) // 整数转字符串拼接
    }
    end_time := time.Now().Unix()
    fmt.Printf("执行耗时 %d秒\n", end_time-start_time)
}

func main() {
    timePrint(time.Now())
    exercises()
}
