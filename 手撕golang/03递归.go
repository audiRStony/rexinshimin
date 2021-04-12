/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 10:04:50
 * @LastEditTime: 2021-04-12 13:06:31
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

//计算n的阶乘
//n < 0 结果错误
//n = 0 结果为1
//n > 0 结果为n! = n * n-1
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
