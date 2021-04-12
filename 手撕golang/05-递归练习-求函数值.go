/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 10:32:35
 * @LastEditTime: 2021-04-12 10:56:22
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */

/*
已知f(1)=3;f(n)= 2*f(n-1)+1；
使用递归的思想编程，求f(n)的值？
*/
package main

import "fmt"

func main() {
	fmt.Println(f(2)) //7
	//fmt.Println(f(3)) //15
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
