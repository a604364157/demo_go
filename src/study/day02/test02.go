//指针
package main

import "fmt"

func func10() {
	a := 10
	fmt.Printf("a的变量地址: %x\n", &a)
}

//指针的声明
func func11() {
	// *T 是指针变量类型,它指向T类型的值
}

func main() {
	func10()
}
