/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 13:19:12
 * @LastEditTime: 2021-04-12 13:37:56
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

/*
go语言提供panic和recover函数用于处理运行时错误，
当调用panic抛出错误时，中断原有的控制流程，常用于不可修复性错误。
recover函数用于终止错误流程，仅在defer语句的函数中有效，用于截取错误流程
recover只能捕捉到最后一个错误
*/

func main() {
	//panic1()
	//success01()
	//failure1()
	failure2()
}
func panic1() {
	defer func() {
		fmt.Println("defer 1")
	}()
	panic("error 00")
	/*
		输出结果：
		defer 1
		panic: error 00
		+ panic msg
	*/
}

//当未发生panic时，recover函数得到的结果为nil
func success01() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("success")
	/*
		输出结果：
		success
		<nil>
	*/
}

//当发生panic则recover函数得到的结果为panic传递的参数
func failure1() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("failure")
	panic("error")
	/*
		result:
		failure
		error
	*/
}

//recover只能捕获到最后一次的panic信息
func failure2() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		fmt.Println("failure1")
		panic("error1")
	}()
	defer func() {
		fmt.Println("failure2")
		panic("error2")
	}()
	/*
		result:
		failure2
		failure1
		error1
	*/
}
