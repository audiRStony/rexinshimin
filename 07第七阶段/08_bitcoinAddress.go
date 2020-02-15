package main

import (
    "golang.org/x/crypto/ripemd160"
    "fmt"
    "crypto/sha256"
    //"bytes"
    "math/big"
    "encoding/hex"
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

func ReverseBytes(data []byte)  {
    for i,j := 0,len(data) -1;i < j; i,j = i+1,j - 1 {
        data[i],data[j] = data[j],data[i]
    }
}


//产生比特币地址
func generateAddress(publickey []byte) []byte {
    //1、产生publickey的hash
    pubkeyHash256 := sha256.Sum256(publickey)
    RIPEMD160Hasher := ripemd160.New()
    _,err := RIPEMD160Hasher.Write(pubkeyHash256[:])

    if err != nil{
        fmt.Println(err)
    }

    pubkeyRIPEMD160 := RIPEMD160Hasher.Sum(nil)

    //2、计算checksum
    versionPayload := append([]byte{0x00},pubkeyRIPEMD160...)
    firstSHA := sha256.Sum256(versionPayload)
    secondSHA := sha256.Sum256(firstSHA[:])
    checksum := secondSHA[:4] //checksum 取前面的4个字节
    //3、base58编码
    fullPayload := append(versionPayload,checksum...)

    //4、比特币地址
    bitcoinAddress := Base58Encode(fullPayload)
    return bitcoinAddress
}

func main() {
    // B7763B66C6D88F941F93F866FE2F3B6911B92CE8359BA1FBFBD6B10A4880D0BA9630BBA53D73C765E47FA2665D4B0A5FC195AD94AFFA88A0F4D68FAD9E94F373
    //外部得到公钥
    publicKey,_ := hex.DecodeString(" B7763B66C6D88F941F93F866FE2F3B6911B92CE8359BA1FBFBD6B10A4880D0BA9630BBA53D73C765E47FA2665D4B0A5FC195AD94AFFA88A0F4D68FAD9E94F373")
    //fmt.Printf("type is %T\n",publicKey)
    bitcoinAddress := generateAddress(publicKey)
    //fmt.Printf("%X\n",bitcoinAddress) //十六进制
    fmt.Printf("%s\n",bitcoinAddress)
}
