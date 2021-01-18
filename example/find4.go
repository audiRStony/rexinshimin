package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
@Time : 2021/1/18 10:40 下午
@Author : audiRS7
@File : find4.go
@Software: GoLand
*/
//文件目录树形结构节点
type dirTreeNode struct {
	name  string
	child []dirTreeNode
}

var iCount int = 0

//递归遍历文件目录
func getDirTree(pathName string) (dirTreeNode, error) {

	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		log.Fatalf("Read dir '%s' failed: %v", pathName, err)
	}
	var tree, childNode dirTreeNode
	tree.name = pathName
	var name, fullName string
	for _, fileDir := range rd {
		name = fileDir.Name()
		fullName = pathName + "/" + name
		if fileDir.IsDir() {
			childNode, err = getDirTree(fullName)
			if err != nil {
				log.Fatalf("Read dir '%s' failed: %v", fullName, err)
			}
		} else {
			childNode.name = name
			childNode.child = nil

			//读取文件内容并打印
			readLine(fullName, processLine)
			iCount++
		}
		tree.child = append(tree.child, childNode)
	}
	return tree, nil
}

//递归打印文件目录
func printDirTree(tree dirTreeNode, prefix string) {
	fmt.Println(prefix + tree.name)
	if len(tree.child) > 0 {
		prefix += "----"
		for _, childNode := range tree.child {
			printDirTree(childNode, prefix)
		}
	}
}

func processLine(line []byte) {
	os.Stdout.Write(line)
}

// read one file and use hookfn to process each line
func readLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

}

func main() {
	dirName := "/Users/jiazhenbing/go/src/example"
	tree, err := getDirTree(dirName)
	if err != nil {
		log.Fatalln("read dir", dirName, "fail: ", err)
	}

	fmt.Printf("\n")
	fmt.Println("File Count:", iCount)

	printDirTree(tree, "")

}
