/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 16:07:03
 * @LastEditTime: 2021-04-12 17:00:27
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */
package main

import "fmt"

//定义user结构体和方法
type User struct {
	name string
	addr string
}

func Newuser(name, tel string) *User {
	return &User{name, tel}
}
func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

//定义employee结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		Newuser(name, addr),
		title,
		salary,
	}
}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)
}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}
func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}
func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {
	me := NewEmployee("Jiaye", "Beijing", "enginer", 8888)
	fmt.Printf("%#v\n", me)
	fmt.Println(me.GetName())      //调用Employee.GetName
	fmt.Println(me.User.GetName()) //调用User.GetName
	me.SetAddr("秦皇岛")
	fmt.Println(me.GetName())
	me.SetSalary(99999)
	fmt.Println(me.GetSalary())
	fmt.Printf("%#v\n", me)
}
