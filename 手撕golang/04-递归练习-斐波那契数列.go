/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 10:21:47
 * @LastEditTime: 2021-04-12 10:30:45
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */

package main

import "fmt"

//请使用递归的方式，求出斐波那契数列1,1,2,3,5,8,13...
//给你一个整数n，求出它的斐波那契数列是多少？
//思路：
//1、当n == 1 || n == 2，返回1
//2、当n>= 2，返回前面两个数的和f(n-1)+f(n-2)

func main() {
	res := fbnq(4)
	fmt.Println(res)

}

func fbnq(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbnq(n-1) + fbnq(n-2)
	}

}
