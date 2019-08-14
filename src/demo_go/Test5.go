package main

import "fmt"

//go语言变量逃逸分析
//go的内存分配是栈，堆

func main() {
	//测试变量逃逸
	var a int
	void()
	fmt.Println(a, dummy(0))
	//go run -gcflags "-m -l"  使用这个命令运行
	//使用 go run 运行程序时，-gcflags 参数是编译参数。其中 -m 表示进行内存分配分析，-l 表示避免程序内联，也就是避免进行程序优化。
	/*
		运行结果如下：
		# command-line-arguments
		./main.go:29:13: a escapes to heap	//变量 a 逃逸到堆
		./main.go:29:22: dummy(0) escapes to heap	//dummy(0)调用逃逸到堆
		./main.go:29:13: main ... argument does not escape
		0 0
	*/

	//测试取地址逃逸
	fmt.Println(dummy2())
	/*
		$ go run -gcflags "-m -l" Test5.go
		# command-line-arguments
		./main.go:15:9: &c escapes to heap
		./main.go:12:6: moved to heap: c
		./main.go:20:19: dummy() escapes to heap
		./main.go:20:13: main ... argument does not escape
		&{}
	*/

	/*
		编译器觉得变量应该分配在堆和栈上的原则是：
		变量是否被取地址。
		变量是否发生逃逸。
	*/
}

//变量和栈的关系
//栈可用于内存分配，栈的分配和回收速度非常快
func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
	//函数执行时，会对 c和x分配内存，并存在栈上
	//函数结束后，保存 c 和 x 的栈内存再出栈释放内存
}

//变量逃逸分析

//本函数测试入口参数和返回值情况
func dummy(b int) int {
	//声明一个c赋值入参并返回
	var c int
	c = b
	return c
}

//空白函数
func void() {

}

//取地址发生逃逸
type Data struct {
}

func dummy2() *Data {
	var c Data
	return &c
}
