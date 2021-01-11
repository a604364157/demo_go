package main

import "fmt"

//生成99乘法表
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
