/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 13:56:19
 * @LastEditTime: 2021-04-12 14:38:07
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  int
	desc string
}

func main() {
	//jiaye()
	//jiaye2()
	jiaye4()

	/*使用new函数进行初始化结构体对象*/
	//var jy3 *Author = new(Author)
	//fmt.Printf("%T: %#v,%#v\n", jy3, jy3, *jy3)
}

//按照结构体字面量初始化结构体值对象
func jiaye() {
	var jy Author = Author{
		1001,
		"JiaYe",
		"BeiJing",
		845090841,
		"少年经不得顺境，中年经不得闲境，晚年经不得逆境", //结尾逗号不能省略
	}
	fmt.Println(jy)
}

//只初始化一部分或者不按照顺序初始化使用命名方式
func jiaye2() {
	jy2 := Author{id: 1002, addr: "ShangHai"}
	fmt.Println(jy2)
}

//使用结构体字面量初始化结构体指针对象
func jiaye4() {
	jiaye4 := &Author{1111, "LuLu", "ShenZhen", 123456, "luluaimeili"}
	fmt.Println(jiaye4)

}
