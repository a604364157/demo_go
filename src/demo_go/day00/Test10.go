package main

import "fmt"

//组合和方法集
//结构类型（struct）为 Go语言提供了强大的类型扩展，
// 主要体现在两个方面：
// 第一，struct 可以嵌入任意其他类型的字段；
// 第二，struct 可以嵌套自身的指针类型的字段。

type X struct {
	a int
}

type Y struct {
	X
	b int
}

type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("X, a = %d\n", x.a)
}

func (y Y) Print() {
	fmt.Printf("y, B = %d\n", y.b)
}

func main() {
	x := X{a: 1}
	y := Y{X: x, b: 2}
	z := Z{Y: y, c: 3}
	println(z.a, z.Y.a, z.Y.X.a)

	z = Z{}
	z.a = 2
	println(z.a, z.Y.a, z.Y.X.a)
}
