package main

/*
@Time : 2020-03-26 18:07
@Author : audiRStony
@File : 12_file_op.go
@Software: GoLand
*/
import (
    "fmt"
    "os"
)

func main() {
    file,err := os.Open("./zzh.txt")
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file.Close()

    var fileTmp [128]byte
    //fileTmp := make([]byte,0,128)
    _,err = file.Read(fileTmp[:])
    if err != nil{
        fmt.Println(err)
        return
    }

    fmt.Println(string(fileTmp[:]))

}

