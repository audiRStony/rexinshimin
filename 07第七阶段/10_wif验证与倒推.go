package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
    var result []byte
    x := big.NewInt(0).SetBytes(input)
    base := big.NewInt(int64(len(b58Alphabet)))
    zero := big.NewInt(0)
    mod := &big.Int{}
    for x.Cmp(zero) != 0 {
        x.DivMod(x, base, mod)
        result = append(result, b58Alphabet[mod.Int64()])
    }
    ReverseBytes(result)
    for _, b := range input {
        if b == 0x00 {
            result = append([]byte{b58Alphabet[0]}, result...)
        } else {
            break
        }
    }
    return result
}
func ReverseBytes(data []byte) {
    for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
}
func generatePrivateKey(hexprivateKey string, compressed bool) []byte {
    versionstr := ""
    if compressed {
        versionstr = "80" + hexprivateKey + "01"
    } else {
        versionstr = "80" + hexprivateKey
    }
    privatekey, _ := hex.DecodeString(versionstr)
    firsthash := sha256.Sum256(privatekey)
    secondhash := sha256.Sum256(firsthash[:])
    checksum := secondhash[:4]
    result := append(privatekey, checksum...)
    base58result := Base58Encode(result)
    return base58result
}
func Base58Decode(input []byte) []byte {
    result := big.NewInt(0)
    zeroBytes := 0
    for _, b := range input {
        if b == '1' {
            zeroBytes++
        } else {
            break
        }
    }
    payload := input[zeroBytes:]
    for _, b := range payload {
        charIndex := bytes.IndexByte(b58Alphabet, b)
        result.Mul(result, big.NewInt(58))
        result.Add(result, big.NewInt(int64(charIndex)))
    }
    decoded := result.Bytes()
    decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
    return decoded
}
func checkWIF(wifprivate string) bool {
    rawdata := []byte(wifprivate)
    base58decodedata := Base58Decode(rawdata)
    fmt.Printf("base58decodedata：%x\n", base58decodedata)
    length := len(base58decodedata)
    if (length < 37) {
        fmt.Printf("长度小于37，一定有问题")
        return false
    }
    private := base58decodedata[:(length - 4)]
    firstsha := sha256.Sum256(private)
    secondsha := sha256.Sum256(firstsha[:])
    checksum := secondsha[:4]
    orignchecksum := base58decodedata[(length - 4):]
    if bytes.Compare(checksum, orignchecksum) == 0 {
        return true
    }
    return false
}
func getPrivateKeyfromWIF(wifprivate string) []byte {
    if checkWIF(wifprivate) {
        rawdata := []byte(wifprivate)
        base58decodedata := Base58Decode(rawdata)
        return base58decodedata[1:33]
    }
    return []byte{}
}
func main() {
    wifkey := generatePrivateKey("d040aafa48e119a1c0902527f4663adf8e90d488f88fdb261a7493688a354547", false)
    fmt.Printf("%s\n", wifkey)
    str := fmt.Sprintf("%s", wifkey)
    privatekey := getPrivateKeyfromWIF(str)
    fmt.Printf("%x\n", privatekey)
}
