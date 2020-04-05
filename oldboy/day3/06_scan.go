package main

/*
@Time : 2020-03-12 09:28
@Author : audiRStony
@File : 06_scan.go
@Software: GoLand
*/

import (
    "fmt"
)
func main() {
    var (
        name   string
        age    int
        gender string
    )

    fmt.Println(name,age,gender)
    //fmt.Scan(&name,&age,&gender) //Scan 扫描，默认以空白为分隔（空格，tab，回车）

    /* 输入时需严格按照格式进行输入：name:LiyaTong age:12 gender:female */
    fmt.Scanf("name:%s age:%d gender:%s\n",&name,&age,&gender) //name 与age之间如有2个空格，扫描输入时也必须有2个空格

    //fmt.Scanln(&name,&age,&gender) //不能使用回车扫描
    fmt.Println(name,age,gender)

}
