package main

/*
@Time : 2020-03-09 19:17
@Author : audiRStony
@File : 02_dispatchCoin.go
@Software: GoLand
*/

import (
    "fmt"
    "strings"
)

/*
共有50枚金币
分配名单：Mattew，Sarah，Augustus，Heidi，Emilie，Peter，Giana，Adriano，Aaron，Elizabeth
规则：
a、名字包含e或者E：1枚金币
b、名字包含i或者I：2枚金币
c、名字包含o或者O：3枚金币
d、名字包含u或者U：4枚金币
计算每个用户分到多少金币，剩余多少金币
*/

var (
    coins        = 50
    users        = []string{"Mattew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
    distribution = make(map[string]int, len(users))
)

// 分发
func dispatchCoin() int {
    sum := 0
    for _, name := range users {
        // strings.Contains方法用于匹配字符串
        if strings.Contains(name, "e") || strings.Contains(name, "E") {
            distribution[name] = distribution[name] + 1
            sum++
            // fmt.Println(" Debug e:",distribution[name])
        }
        if strings.Contains(name, "i") || strings.Contains(name, "I") {
            distribution[name] = distribution[name] + 2
            sum = sum + 2
        }
        if strings.Contains(name, "o") || strings.Contains(name, "O") {
            distribution[name] = distribution[name] + 3
            sum = sum + 3
        }
        if strings.Contains(name, "u") || strings.Contains(name, "U") {
            distribution[name] = distribution[name] + 4
            sum = sum + 4
        }
    }
    return coins - sum
}

func main() {
    surplus := dispatchCoin()
    fmt.Println("剩余:", surplus)

    for k, v := range distribution {
        fmt.Println(k, v)
    }
}
