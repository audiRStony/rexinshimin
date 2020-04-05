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

    var fileTmp [128]byte //func (f *File) Read(b []byte) (n int, err error)
    //fileTmp := make([]byte,0,128)
    _,err = file.Read(fileTmp[:])//将读取的内容存到[]byte中。返回的int类型为字节数
    if err == io.EOF{ //判断是否读取到文件末尾
        fmt.Println(文件读取完毕)
        return
    }

    fmt.Println(string(fileTmp[:]))//将缓存的[]byte转化为字符串

}

