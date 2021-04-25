package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

/*
 @Author: JiaYe
 @Date: 2021/4/22 8:18 下午
 @SoftWare: Goland
*/

const MB = 1024 * 1024

func main() {
	blocks := make([][MB]byte, 0)
	fmt.Println("child pid is", os.Getpid())
	for i := 0; ; i++ {
		blocks = append(blocks, [MB]byte{})
		printMemUsage()
		time.Sleep(time.Second)
	}
}
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tSys = %v Mib \n", bToMb(m.Sys))
}

func bToMb(b uint64) uint64 {
	return b / MB
}
