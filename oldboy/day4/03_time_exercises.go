package main

/*
@Time : 2020-03-18 00:37
@Author : audiRStony
@File : 03_time_exercises.go
@Software: GoLand
*/
import (
    "fmt"
    "strconv"
    "time"
)

// 格式化打印
func timePrint(t time.Time) {
    timer := t.Format("2006/01/02 15:04:05")
    fmt.Println(timer)
}

// 测试用例
func exercises() {
    start_time := time.Now()
    // fmt.Println(start_time)
    str := "" // 声明一个字符串
    for i := 0; i < 100000; i++ { // for循环10W次拼接
        str += "golang" + strconv.Itoa(i) // 整数转字符串拼接
    }
    time.Sleep(time.Millisecond * 30)
    // fmt.Printf("执行耗时 %dS\n", time.Since(start_time)/1000/1000/1000)
    fmt.Println("执行耗时:", time.Since(start_time))
}

func main() {
    // timePrint(time.Now())
    exercises()
}
