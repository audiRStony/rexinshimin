package main

import (
    "math/big"
    "fmt"
    "bytes"
)

//定义切片存储base58字母
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte  {
    var result  []byte //定义一个切片。返回值
    x := big.NewInt(0).SetBytes(input) //把字节数组input转化为大整数bigint
    base := big.NewInt(int64(len(b58Alphabet))) //长度58的大整数
    zero := big.NewInt(0) //0的大整数

    mod := &big.Int{} //大整数的指针
    //循环。不停地对58取余数，大小为58
    for x.Cmp(zero) != 0 {
        x.DivMod(x,base,mod) //对x取余数

        result = append(result,b58Alphabet[mod.Int64()]) //将余数添加到数组当中
    }

    ReverseBytes(result) //反转字节数组
    //如果字节数组的前位为字节0，替换为1。比特币中特殊做法
    for _,b := range input{
        if b ==0x00{
            result = append([]byte{b58Alphabet[0]},result...)
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
        charIndex := bytes.IndexByte(b58Alphabet,b) //反推出余数
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

func main() {

    /* 测试反转
    org := []byte("qwerty")
    fmt.Println(string(org))
    ReverseBytes(org)
    fmt.Println(string(org))*/

    fmt.Printf("%s\n",string(Base58Encode([]byte("hello audiRStony"))))
    fmt.Printf("%s\n",string(Base58decode([]byte("DthHcFYf1DPUyatoLuZm68"))))
}
