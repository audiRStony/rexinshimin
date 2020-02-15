package main

import (
    "encoding/hex"
    "crypto/sha256"
    "math/big"
    "bytes"
    "fmt"
)

var b58Alphaet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte  {
    var result   []byte //定义一个切片。返回值
    x := big.NewInt(0).SetBytes(input) //把字节数组input转化为大整数bigint
    base := big.NewInt(int64(len(b58Alphaet))) //长度58的大整数
    zero := big.NewInt(0) //0的大整数

    mod := &big.Int{} //大整数的指针
    //循环。不停地对58取余数，大小为58
    for x.Cmp(zero) != 0 {
        x.DivMod(x,base,mod) //对x取余数

        result = append(result,b58Alphaet[mod.Int64()]) //将余数添加到数组当中
    }

    ReverseBytes(result) //反转字节数组
    //如果字节数组的前位为字节0，替换为1。比特币中特殊做法
    for b := range input{
        if b ==0x00{
            result = append([]byte{b58Alphaet[0]},result...)
        }else {
            break
        }
    }

    return  result
}

func Base58decode(input []byte) []byte  {

    result := big.NewInt(0)
    zeroBytes := 0

    for _,b := range input {
        if b == '1' {
            zeroBytes++
        }else {
            break
        }
    }

    //除去前面的1
    payload := input[zeroBytes:]
    //循环，逆推结果
    for _,b := range payload{
        charIndex := bytes.IndexByte(b58Alphaet,b) //反推出余数
        result.Mul(result,big.NewInt(58)) //之前的结果乘以58
        result.Add(result,big.NewInt(int64(charIndex))) //加上这个余数
    }

    decoded := result.Bytes() //大整数转化为字节数组
    decoded = append(bytes.Repeat([]byte{0x00},zeroBytes),decoded...)

    return decoded
}


//反转字节数组
func ReverseBytes(data []byte)  {
    for i,j := 0,len(data) -1;i < j; i,j = i+1,j - 1 {
        data[i],data[j] = data[j],data[i]
    }
}


func generatePrivateKey(hexprivateKey string,compressed bool) []byte  {
    versionstr := ""

    //判断是否对应的是压缩的公钥，如果是，需要在前面加上0x01这个字节，同时任意的私钥，需要在前面加上0x80的字节
    if compressed{
        versionstr = "80" + hexprivateKey + "01"
    }else {
        versionstr = "80" + hexprivateKey
    }

    //字符串转化为16进制字节
    privatekey,_ := hex.DecodeString(versionstr)

    //通过double hash，计算checksum的值,checksum 是两次sha256以后的前4个字节
    firsthash := sha256.Sum256(privatekey)
    secondhash := sha256.Sum256(firsthash[:])
    checksum := secondhash[:4]

    //拼接
    result := append(privatekey,checksum...)
    //通过base58编码
    base58result := Base58Encode(result)
    return base58result

}

func main() {
    wifkey := generatePrivateKey("f59a72636fb1eb79a9c31a078992fc8812d777593bb06743eba32666f99906f9",true)
    fmt.Printf("%s\n",wifkey)
}
