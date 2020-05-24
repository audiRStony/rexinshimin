package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
@Time : 2020/5/24 3:41 下午
@Author : audiRS7
@File : 10_syncMap
@Software: GoLand
*/

/*错误示例*/
//fatal error: concurrent map writes
//golang中的map不是并发安全的，某个goroutine在set map的时候，其他goroutine可能就去读了，导致混乱
//var m = make(map[string]int)
/*func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}*/
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n) //将int类型转化为string类型
//			set(key, n)
//			fmt.Printf("key = %s,value = %d\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

/*sync.Map示例*/
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("key = %s,value = %d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
