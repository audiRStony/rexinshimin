package main

/*
@Time : 2020-03-27 01:06
@Author : audiRStony
@File : 13_file_buffio.go
@Software: GoLand
*/
import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func main() {
    file,err := os.Open("./zzh.txt")
    if err != nil{
        fmt.Println("打开文件失败")
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for   {
        bufstr,err := reader.ReadString('\n') /*读到指定字符停止 单引号为字符*/
        if err != nil {
            if err == io.EOF{ //读到文件结尾
                os.Exit(0)
            }
            fmt.Println("文件出错")
        }
        fmt.Print(bufstr)
    }
}

