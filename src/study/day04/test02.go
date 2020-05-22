//通道channel 不要通过共享内存来通信，而应该通过通信来共享内存
//通道是什么，通道就是goroutine之间的通道。它可以让goroutine之间相互通信
package main

import "fmt"

func func11() {
	//通道得声明
	a := make(chan int)
	fmt.Println(a)
	//channel的数据类型
	fmt.Printf("%T,%p\n", a, a)
}

func main() {
	func11()
}
