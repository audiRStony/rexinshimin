package main

import (
    "fmt"
    "crypto/sha256"
    "os"
    "io"
)

func main() {
    /*//hash函数第一种
    sum := sha256.Sum256([]byte("audiRStony"))
    fmt.Printf("%X\n",sum)*/

    //第二种
  /*  h := sha256.New()
    h.Write([]byte("audiRStony"))
    fmt.Printf("%X\n",h.Sum(nil))*/

    //第三种。文件操作

    h := sha256.New()

    file_data,err1 := os.Open("test_hash.txt")
    defer file_data.Close()

    if err1 != nil{
        fmt.Println(err1)
    }

    if _,err2 := io.Copy(h,file_data); err2 != nil{
        fmt.Println(err2)
    }
    fmt.Printf("%X\n",h.Sum(nil))

}
