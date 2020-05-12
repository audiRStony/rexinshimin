package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
@Time : 2020/5/12 10:47 下午
@Author : audiRS7
@File : 05_ProducerCustomer
@Software: GoLand
*/

//生产者消费者模型
/*使用goroutine和channel实现一个简单的生产者消费者模型
生产者：产生随机数
消费者：计算每个随机数的每位数的和，如 123=6
1个生产者 20个消费者*/

var itemChan chan *item
var resultChan chan *result

type item struct {
	id  int64
	num int64
}
type result struct {
	item *item
	sum  int64
}

//producer
func producer(ch chan *item) {
	//1、生产随机数
	var id int64
	for {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}
		//2、把随机数发送到管道
		ch <- tmp
	}
}

//计算器。计算每位数字的和
func calc(num int64) int64 {
	var sum int64
	for num > 0 { //0
		sum = sum + num%10 //sum = 5 + 1
		num = num / 10     //num = 0
	}
	return sum
}

//customer
func customer(ch chan *item, resultChan chan *result) {
	for tmp := range ch {
		sum := calc(tmp.num)
		//构造result结构体
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}
}

func startworker(n int, ch chan *item, resultChan chan *result) {
	for i := 1; i <= n; i++ {
		go customer(ch, resultChan)
	}
}

//printer
func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id:%v,num:%v,sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}
}

func main() {
	itemChan := make(chan *item, 100)
	resultChan := make(chan *result, 100)
	go producer(itemChan)
	//go customer()
	startworker(20, itemChan, resultChan)
	/*rand.Seed(time.Now().UnixNano()) //设置种子，每次执行都能产生随机数
	ret1 := rand.Intn(1, 101)        //左包含右不包含
	fmt.Println(ret1)
	ret2 := rand.Int63() //int64的正整数
	fmt.Println(ret2)*/

	//打印结果
	printResult(resultChan)
}
