/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 15:30:12
 * @LastEditTime: 2021-04-12 16:03:44
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

//定义结构体Dog
type Dog struct {
	name string
}

func (dog Dog) call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

//为结构体Dog定义方法Setname
//结构体为值传递，修改数据需使用指针
func (dog *Dog) Setname(name string) {
	dog.name = name

}

//调用
/*
调用方法通过自定义类型的对象.方法名进行调用，在调用的过程中对象传递(赋值)给方法的接受者(值类型，拷贝)
*/
func main() {
	dog := Dog{"dahuang"}
	dog.call()
	dog.Setname("mingshiyin")
	dog.call()
}
