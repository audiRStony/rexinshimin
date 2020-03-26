package main

/*
@Time : 2020-03-25 16:52
@Author : audiRStony
@File : 10_carInterface.go
@Software: GoLand
*/
import (
	"fmt"
)

/*定义具有价值的不同事物的"合同" */
type Valuable interface {
	GetValue() float32
}
type StockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

// 获取stock的值(价格)
func (s StockPosition) GetValue() float32 {
	return s.sharePrice * s.count
}

type Car struct {
	make  string
	model string
	price float32
}

// 获取car的值(价格)
func (c Car) GetValue() float32 {
	return c.price
}

func ShowValue(asset Valuable) {
	fmt.Printf("资产的价值是：%f\n", asset.GetValue())
}

func main() {
	var o Valuable = StockPosition{"GOOD", 577.20, 4}
	ShowValue(o)
	o = Car{"BMW", "M3", 66500}
	ShowValue(o)
}
