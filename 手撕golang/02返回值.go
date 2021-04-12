/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-09 17:47:13
 * @LastEditTime: 2021-04-12 10:01:04
 * @LastEditors: JiaYe
 */
package main

import "fmt"

func main() {
	//fmt.Println(calc(10, 5))
	//sayHello()
	//sayHello1("JiaYe")
	fmt.Println(add(1, 2))
}

func calc(n1, n2 int) (int, int, int, int) {

	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

//无参无返回值
func sayHello() {
	fmt.Println("hello goland")
}

//有参无返回值
func sayHello1(name string) {
	//fmt.Printf("Hi, %v\n", name)
	fmt.Println("Hi,", name)
}

//有参有返回值
func add(n1, n2 int) int {
	return n1 + n2
}
