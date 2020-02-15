package main

import (
    "fmt"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "math/big"
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

    hash := sha256.Sum256([]byte("hello tony"))

    r,s,err := ecdsa.Sign(rand.Reader,&privatekey,hash[:])

    if err!= nil{
        fmt.Println(err)
    }

    /*//生成数据签名
    signature := append(r.Bytes(),s.Bytes()...)
    fmt.Printf("%X\n",signature)*/

    //数据签名验证
    curve := elliptic.P256()
    keylength := len(publickey)

    x := big.Int{}
    y := big.Int{}
    x.SetBytes(publickey[:(keylength/2)])
    y.SetBytes(publickey[(keylength/2):])

    rawPublic := ecdsa.PublicKey{curve,&x,&y}
    if ecdsa.Verify(&rawPublic,hash[:],r,s) == false{
        fmt.Println("验证失败")
    }else {
        fmt.Println("验证成功")
    }


//r,s接收的是hash的坐标轴，所以验证hash2时失败
    hash2 := sha256.Sum256([]byte("hello nike"))
    if ecdsa.Verify(&rawPublic,hash2[:],r,s) == false{
        fmt.Println("验证失败")
    }else {
        fmt.Println("验证成功")
    }



}

