package main

import (
    "fmt"
    "math/big"
    "bytes"
)

var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {

    var result []byte

    x := big.NewInt(0).SetBytes(input)

    base := big.NewInt(int64(len(base58Alphabet)))

    zero := big.NewInt(0)

    mod := &big.Int{}

    for x.Cmp(zero) != 0 {

        //fmt.Println(x)

        //fmt.Println(base)

        x, mod = x.DivMod(x, base, mod) //DIVMod,x除以base，返回商和余数

        result = append(result, base58Alphabet[mod.Int64()])//余数去对应字母表

        //fmt.Println(mod)

    }

    if input[0] == 0x00 {

        result = append(result, base58Alphabet[0])

    }

    ReverseBytes(result)

    return result

}

func Base58Decode(input []byte) []byte {

    result := big.NewInt(0)

    for _, b := range input {

        charIndex := bytes.IndexByte(base58Alphabet, b)

        result.Mul(result, big.NewInt(58))

        result.Add(result, big.NewInt(int64(charIndex)))

        //fmt.Println(result)

    }

    decoded := result.Bytes()

    if input[0] == base58Alphabet[0] {

        decoded = append([]byte{0x00}, decoded...)

    }

    return decoded

}

func ReverseBytes(data []byte) {

    for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {

        data[i], data[j] = data[j], data[i]

    }

}

func main() {

    fmt.Printf("%x\n", Base58Encode([]byte("abcdefg")))
    fmt.Printf("%x\n", Base58Decode(Base58Encode([]byte("abcdefg"))))

}
