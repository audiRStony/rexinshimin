package main

/*
@Time : 2020-03-27 02:02
@Author : audiRStony
@File : 14_ioutil.go
@Software: GoLand
*/

import (
    "io/ioutil"
    "io"
    "os"
    "fmt"
)

func main() {
    content,err := ioutil.ReadFile("./zzh.txt") //返回[]byte，error
    if err !=nil{
        if err == io.EOF{
            os.Exit(0)
        }
        return
    }
    fmt.Println(string(content))
}

