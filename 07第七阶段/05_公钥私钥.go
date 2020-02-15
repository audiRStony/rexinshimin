package main

import (
    "fmt"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
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
    privatekey,publickey := NewKeypair()
    //fmt.Printf("public key: %X\n",publickey)

    //打印私钥,曲线上的x点
    fmt.Printf("%x\n",privatekey.D.Bytes())
    //打印公钥，曲线上的y点
    fmt.Printf("%x\n",publickey)
}
