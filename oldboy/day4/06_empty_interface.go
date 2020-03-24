package main

/*
@Time : 2020-03-24 17:12
@Author : audiRStony
@File : 06_empty_interface.go
@Software: GoLand
*/

import (
    "fmt"
    "time"
)

func showType(a interface{}) {
    fmt.Printf("Type:%T\n", a)
}

func main() {
    /*var x interface{} //空接口.传入任何参数都满足接口条件
      x = 100
      showType(x)

      x = false
      showType(x)

      x = time.Second
      showType(x)
    */
    // 声明一个值为空接口的map
    var StuInfo = make(map[string]interface{}, 100)
    StuInfo["阿珂"] = 100
    StuInfo["孙尚香"] = true
    StuInfo["白起"] = 66.66
    StuInfo["亚瑟"] = "对抗路"
    StuInfo["赵云"] = time.Now()
    fmt.Println(StuInfo)
}
