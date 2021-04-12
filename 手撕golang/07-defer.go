/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 13:08:03
 * @LastEditTime: 2021-04-12 13:16:33
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

func main() {
	/*
		多个defer，先defer的最后被执行
	*/
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	fmt.Println("over")
	/*
		输出结果
		over
		defer 3
		defer 2
		defer 1
	*/
}
