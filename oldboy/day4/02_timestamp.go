package main

/*
@Time : 2020-03-17 20:50
@Author : audiRStony
@File : 02_timestamp.go
@Software: GoLand
*/

import (
    "fmt"
    "time"
)
//时间戳
func timestampDemo(timestamp int64) {
    timeObj := time.Unix(timestamp, 0) // 将时间戳转化为时间格式
    fmt.Println(timeObj)
    year := timeObj.Year()
    month := timeObj.Month()
    day := timeObj.Day()
    // day1 := time.Unix().Day()
    hour := timeObj.Hour()
    minute := timeObj.Minute()
    second := timeObj.Second()
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//日期格式化
func formatDemo()  {
    now := time.Now()
    //格式化的模板为go语言诞生的时间：2006年1月2日15时04分05秒
    fmt.Println(now.Format("2006-01-02 15:04:05"))
    fmt.Println(now.Format("2006/01/02 15:04:05"))
    fmt.Println(now.Format("15:04:05 2006/01/02"))
    fmt.Println(now.Format("2006/01/02"))
    fmt.Println(now.Format("2006-01-02"))
}
func main() {
    /*now := time.Now()
    tm := now.Unix() + 10000
    timestampDemo(tm)*/
    formatDemo()
}
