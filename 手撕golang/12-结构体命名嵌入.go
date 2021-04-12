/*
 * @Description:
 * @Author: JiaYe
 * @Date: 2021-04-12 14:54:03
 * @LastEditTime: 2021-04-12 15:19:37
 * @LastEditors: JiaYe
 * @Descripttion:
 * @version:
 */

/*
结构体命名嵌入是指结构体中的属性对应的类型也是结构体
*/

package main

import "fmt"

func main() {

	//定义地址的结构体(区域、街道、门牌号)
	type address struct {
		region string
		street string
		no     int
	}

	//定义收货人的结构体(名字、电话、地址)
	type user1 struct {
		name string
		tel  int
		addr address
	}
	//初始化&声明
	milaidi := user1{name: "milaidi", tel: 123456, addr: address{region: "ChaoYang", street: "huixinxijie", no: 3}}
	fmt.Println(milaidi)
	milaidi.addr.region = "guomao"
	//fmt.Println(milaidi.addr.region) //修改
}
