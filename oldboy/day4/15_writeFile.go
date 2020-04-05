package main

/*
@Time : 2020-03-27 03:16
@Author : audiRStony
@File : 15_writeFile.go
@Software: GoLand
*/
import (
    "os"
    "fmt"
    "bufio"
    "io/ioutil"
)

func main() {
    //ioWriter()
    //bufioWriter()
    ioutilWriter()
}

func ioWriter()  {
    file,err := os.OpenFile("./abc.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)//不使用append则在文件开头加入写的内容
    if err != nil{
        fmt.Println("open file failed",err)
        return
    }
    defer file.Close()

    content_str := "LiyaTong\n"
    file.Write([]byte("娜扎\n")) //写入字节切片
    file.WriteString(content_str) //写入string
}

//bufio方式
func bufioWriter()  {
    file,err := os.OpenFile("./bcd.txt",os.O_CREATE|os.O_CREATE|os.O_APPEND,0644)
    if err != nil{
        fmt.Println("open file failed",err)
        return
    }
    defer file.Close()

    writer1 := bufio.NewWriter(file)
    for i:=0;i<10 ;i++  {
        writer1.WriteString("hello 娜扎\n")
    }
    writer1.Flush() //将缓存中的内容写入文件
}

//ioutil方式
func ioutilWriter()  {
    contentStr := "hello LiyaTong\n"
    err := ioutil.WriteFile("./LiyaTong.txt",[]byte(contentStr),0644)
    if err !=nil {
        fmt.Println("write file failed")
        return
    }
}
