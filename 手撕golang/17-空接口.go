/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-19 16:08:28
 * @LastEditTime: 2021-04-19 16:41:37
 * @LastEditors: JiaYe
 */
package main

import "fmt"

/*
不包含任何函数签名的接口叫做空接口，空接口声明的变量可以赋值为任何类型的变量(任意接口)
*/
func main() {
	var i01 interface{} = 1
	var i02 interface{} = true
	var i03 interface{} = &i01
	var i04 interface{} = "我是露露"
	var i05 interface{} = [...]int{1, 2, 3, 4, 5}
	var i06 interface{} = []int{10, 11, 12}
	var i07 interface{} = map[string]string{"name": "age"}
	var i08 interface{} = struct{ x, y int }{1, 2}
	/*fmt.Printf("%#v\n", i01)
	fmt.Printf("%#v\n", i02)
	fmt.Printf("%#v\n", i03)
	fmt.Printf("%#v\n", i04)
	fmt.Printf("%#v\n", i05)
	fmt.Printf("%#v\n", i06)
	fmt.Printf("%#v\n", i07)
	fmt.Printf("%#v\n", i08)
	*/

	printType(i01)
	printType(i02)
	printType(i03)
	printType(i04)
	printType(i05)
	printType(i06)
	printType(i07)
	printType(i08)
}

//使用场景：常声明函数参数类型为interface{}，用于接收任意类型的变量

func printType(vs ...interface{}) {
	for _, data := range vs {
		switch data.(type) {
		case nil:
			fmt.Println("nil")
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		case [5]int:
			fmt.Println("[5]int")
		case []int:
			fmt.Println("[]int")
		case map[string]string:
			fmt.Println("map[string]string")
		default:
			fmt.Println("unknow")
		}
	}
}
