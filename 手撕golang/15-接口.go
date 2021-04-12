/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 17:15:05
 * @LastEditTime: 2021-04-12 17:45:17
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */

package main

import "fmt"

type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

//赋值
/*
当自定义类型实现了接口类型中声明的所有函数时，则该类型的对象可以赋值给接口变量
并使用接口变量调用实现的接口
*/
/*a 方法接受者全为值类型的方法*/

//定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

func NewEmailSender(addr string, port int, user, password string) EmailSender {
	return EmailSender{addr, port, user, password}
}

//接受者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给: %s,内容为: %s\n", to, msg)
	return nil
}

//接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给: %s,内容为: %s\n", to, msg)
	}
	return nil
}

func main() {
	//赋值Sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "lulu", "123456"}
	//fmt.Printf("%T: %v\n", sender01, sender01)
	sender01.Send("audirstony@gmail.com", "Hi golang")
	sender01.SendAll([]string{"jzb0424@163.com", "jzb89757@sina.com"}, "你好 世界")
}
