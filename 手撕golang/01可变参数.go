package main

import "fmt"

func main() {
	//PrintArgs(1, 2)
	//PrintArgs(1, 2, "3", "4")
	PrintArgs(1, 2, []string{"3", "4", "5"}...)
	args := []string{"6", "7", "8", "9"}
	PrintArgs(1, 2, args...)
	PrintArgs(1, 2, args[:3]...) //左包含右不包含

}

func PrintArgs(n1, n2 int, args ...string) {
	//fmt.Printf("%T %T %T\n", n1, n2, args)
	fmt.Println(n1, n2, args)
}
