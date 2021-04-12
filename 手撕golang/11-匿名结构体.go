/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 14:41:04
 * @LastEditTime: 2021-04-12 14:52:22
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

/*
   在定义变量时将类型指定为结构体的时候，此时叫做匿名结构体。匿名结构体常用于初始化
   一次结构体变量的场景
*/

//匿名结构体
var coordinate struct {
	longitude float64
	latitude  float64
}

func main() {

	me := struct {
		name string
		age  int
		addr string
	}{"sunshangxiang", 26, "shuguo"}

	//匿名结构体数组
	studs := []struct {
		name string
		addr string
	}{{"makeboluo", "yidali"}, {"zhaoyun", "shuguo"}}

	fmt.Println(coordinate, coordinate.longitude, coordinate.latitude)
	fmt.Println(me.name, me.age, me.addr)
	fmt.Println(studs)
}
