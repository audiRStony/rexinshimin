/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 11:31:37
 * @LastEditTime: 2021-04-12 11:47:48
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

func main() {
	name, age, heigh, isBoy := "JiaYe", 30, 1.77, false //定义字符串、数值、布尔类型
	pointer := new(int)                                 //定义指针类型
	scores := [...]int{1, 2, 3}                         //定义数组类型
	names := make([]string, 1, 3)                       //定义切片
	users := make(map[int]string)                       //定义映射

	name2, age2, heigh2, isBoy2, pointer2, scores2, names2, users2 := name, age, heigh, isBoy, pointer, scores, names, users
	//测试
	name2 = "LuLu"
	age2 = 27
	scores[0] = 0
	names2[0] = "LuLu"
	//pointer2 = &pointer
	users2[1] = "LuLu"

	fmt.Println(name, age, heigh, isBoy)
	fmt.Println(name2, age2, heigh2, isBoy2, pointer2, scores2, names2, pointer2, users2)
}
