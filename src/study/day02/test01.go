//函数
package main

import (
	"fmt"
	"math"
)

//可变参
func func01(arg ...string) {
	fmt.Println(arg)
}

//参数的值传递和引用传递
func func02() {
	//声明函数变量(开方)[值传递示例]
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//[引用传递示例](指针的类型)
	add := func(a *int) int {
		*a = *a + 1
		return *a
	}
	a := 1
	fmt.Println(a)
	b := add(&a)
	fmt.Println(a, b)
}

//函数的返回值
func func03() {
	//函数的返回值定义在函数后,返回值可以有多个
	a := func(x, y string) (string, string) {
		return y, x
	}
	fmt.Println(a("1", "2"))
}

//_空白标识符(空白标识符一般用于占位)
func func04() {
	a := func(x, y string) (string, string) {
		return y, x
	}
	c, _ := a("1", "2")
	fmt.Println(c)
}

// defer延迟函数
func func05() {
	a, b := 1, 2
	defer fmt.Println(a)
	fmt.Println(b)
}

func main() {
	//func01("a")
	//func01("a", "b")
	//func02()
	//func03()
	//func04()
	func05()
}
