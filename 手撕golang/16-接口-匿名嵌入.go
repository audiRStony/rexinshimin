/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-19 14:54:28
 * @LastEditTime: 2021-04-19 16:21:24
 * @LastEditors: JiaYe
 */
package main

import "fmt"

//定义Sender接口
type Sender1 interface {
	Send1(msg string) error
}

//定义接收接口
type Receiver interface {
	Receive() (string, error)
}

//定义Connection接口，由Sender和Recevier组成
type Connection interface {
	Sender1
	Receiver
	Open() error
	Close() error
}

//定义TCPConnection结构体，并实现Connection接口中open、send、receiver、close方法
type TCPConnection struct {
	addr string
	port int
}

func (conn TCPConnection) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnection) Send1(msg string) error {
	fmt.Println("发送消息:", msg)
	return nil
}

func (conn TCPConnection) Receive() (string, error) {
	fmt.Println("接收成功")
	return "", nil
}

func (conn TCPConnection) Close() error {
	fmt.Println("关闭")
	return nil
}

func main() {
	var conn Connection = TCPConnection{"127.0.0.1", 8888}
	conn.Open()
	conn.Send1("你好 赵云")
	msg, _ := conn.Receive()
	fmt.Printf("%q\n", msg)
	conn.Close()
}
