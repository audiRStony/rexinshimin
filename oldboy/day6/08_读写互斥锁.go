package main

import (
	//"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"fmt"
	"sync"
	"time"
)

/*
@Time : 2020/5/24 1:48 下午
@Author : audiRS7
@File : 08_读写互斥锁
@Software: GoLand
*/
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex   //互斥锁
	rwlock sync.RWMutex //读写互斥锁:多个goroutine同时加读锁，写的时候加写锁
)

/*
读的次数远远大于写的次数，读写互斥锁的执行效率高，如读写次数相差无几，读写互斥锁和互斥锁的执行效率差不多
*/

func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock() //读锁
	x = x + 1
	//lock.Unlock()//解锁
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()

}

func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	//fmt.Println(x)
	time.Sleep(time.Millisecond * 1)
	rwlock.RUnlock()
	//lock.Unlock()

}

func main() {
	startTime := time.Now()
	//写10次
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	//读1000次
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go read()
	}
	endTime := time.Now()
	fmt.Printf("Used %v \n", endTime.Sub(startTime))

}
