package main

import (
    "fmt"
    "encoding/hex"
    "crypto/sha256"
)

func ReverseBytes(data []byte)  {
    for i,j := 0,len(data) -1;i < j; i,j = i+1,j - 1 {
        data[i],data[j] = data[j],data[i]
    }
}

func main() {
    //字符串hash转化为字节hash
    hash1,_ := hex.DecodeString("16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f")
    hash2,_ := hex.DecodeString("cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061")

    //大小端转化
    ReverseBytes(hash1)
    ReverseBytes(hash2)

    //拼接
    rawdata := append(hash1,hash2...)
    //double sha256
    firsthash := sha256.Sum256(rawdata)
    secondhash := sha256.Sum256(firsthash[:])
    //生成默克尔根并反转，与浏览器大端匹配
    merkelRoot := secondhash[:]
    ReverseBytes(merkelRoot)

    fmt.Printf("%x\n",merkelRoot)


}
