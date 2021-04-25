package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

/*
 @Author: JiaYe
 @Date: 2021/4/23 11:44 下午
 @SoftWare: Goland
 @Description:
	1、通过配置参数memory，修改memory.limit_in_bytes和memory.swappiness来设置最大内存使用量
	2、通过cmd.Start()启动一个进程
	3、将新生成的进程号cmd.Process.Pid写到cgroup.procs
	4、通过cmd.Wait()接收命令输出结果
	5、如果返回结果为kill信号的时候，能够重启任务
*/

var (
	rssLimit   int
	cgroupRoot string
)

const (
	procsFile       = "cgroup.procs"
	memoryLimitFile = "memory.limit_in_bytes"
	swapLimitFile   = "memory.swappiness"
)

func init() {

	flag.IntVar(&rssLimit, "memory", 10, "memory limit with MB.")
	flag.StringVar(&cgroupRoot, "root", "/sys/fs/cgroup/memory/climits", "cgroup root path")

}

func main() {
	flag.Parse()
	//set memory limit
	mPath := filepath.Join(cgroupRoot, memoryLimitFile)
	whiteFile(mPath, rssLimit*1024*1024)

	//set swap memory limit to zero
	sPath := filepath.Join(cgroupRoot, swapLimitFile)
	whiteFile(sPath, 0)

	go startCmd("./simpleapp")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("Got signal:", s)

}

func whiteFile(path string, value int) {
	if err := ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", value)), 0755); err != nil {
		log.Panic(err)
	}
}

type ExitStatus struct {
	Signal os.Signal
	Code   int
}

func startCmd(command string) {
	restart := make(chan ExitStatus, 1)
	runner := func() {
		cmd := exec.Cmd{
			Path: command,
		}
		cmd.Stdout = os.Stdout
		//start app
		if err := cmd.Start(); err != nil {
			log.Panic(err)
		}
		fmt.Println("app pid", cmd.Process.Pid, "to file cgroup.proces")
		//set cgroup proces id
		pPath := filepath.Join(cgroupRoot, procsFile)
		whiteFile(pPath, cmd.Process.Pid)

		if err := cmd.Wait(); err != nil {
			fmt.Println("cmd return with error:", err)
		}
		status := cmd.ProcessState.Sys().(syscall.WaitStatus)

		options := ExitStatus{
			Code: status.ExitStatus(),
		}

		if status.Signaled() {
			options.Signal = status.Signal()
		}
		cmd.Process.Kill()
		restart <- options
	}
	go runner()

	for {
		status := <-restart
		switch status.Signal {
		case os.Kill:
			fmt.Println("app is killed by system")
		default:
			fmt.Println("app exit whit code:", status.Code)
			return
		}
		fmt.Println("restart app...")
		go runner()
	}
}
