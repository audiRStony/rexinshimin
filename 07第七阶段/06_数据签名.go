package main

import (
    "fmt"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
)

func NewKeypair() (ecdsa.PrivateKey,[]byte)  {
    //生成椭圆曲线。secp256r1曲线 （比特币当中的曲线是secp256k1）
    curve := elliptic.P256()

    private,err1 := ecdsa.GenerateKey(curve,rand.Reader)
    if err1 != nil{
        fmt.Println(err1)
    }

    pubkey := append(private.PublicKey.X.Bytes(),private.PublicKey.Y.Bytes()...)
    return *private,pubkey
}


func main() {
    //调用函数生成曲线
    privatekey,_ := NewKeypair()

    hash := sha256.Sum256([]byte("hello tony"))

    r,s,err := ecdsa.Sign(rand.Reader,&privatekey,hash[:]) //rand.Reader生产种子随机数

    if err!= nil{
        fmt.Println(err)
    }

    //生成数据签名
    signature := append(r.Bytes(),s.Bytes()...)
    fmt.Printf("%X\n",signature)//返回的数据签名为椭圆曲线上的随机数


}

