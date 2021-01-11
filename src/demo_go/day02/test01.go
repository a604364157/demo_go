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

// defer
func func05() {
	//延迟函数
	a, b := 1, 2
	defer fmt.Println(a)
	fmt.Println(b)
}

func func06() {
	//延迟参数(也就是延迟函数所接收的参数是在其被读取到时的参数,这很好理解,函数调用的参数,会在函数上复制)
	c := 5
	defer fmt.Println("延迟后c=", c)
	c = 10
	fmt.Println("c=", c)
}

func func07() {
	//堆栈的延迟(可以发现多个延迟是先进后出的栈结构)
	a := "123456"
	fmt.Println("a=", a)
	for _, v := range a {
		defer fmt.Print(v)
	}
}

//在go中,函数本质也是一种数据类型,可以看作复合数据类型
//go是基于c/c++升级,在函数定义上类似于c/c++
func func08() {
	//匿名函数
	func() {
		fmt.Println("这是一个匿名函数")
	}()
	//带参匿名函数
	func(a, b int) {
		fmt.Println("带参匿名函数", a, b)
	}(1, 2)
	//带入参出参的匿名函数
	inner := func(a, b int) int {
		return a * b
	}
	fmt.Println(inner(2, 3))
	//回调函数(函数可以作为参数传递)
	add := func(a, b int) int {
		return a + b
	}
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println("------------调用高阶函数oper------------")
	fmt.Println(oper(1, 2, add))
	fmt.Println(oper(2, 1, sub))
}

func oper(a, b int, fun func(int, int) int) int {
	//可以可以看到fun是指向一个引用地址,及传入函数的地址
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}

//闭包(go支持将一个函数作为参数,也可以将函数作为返回值)
//一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量
//(外层函数中的参数，或者外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数
//这个内层函数和外层函数的局部变量，统称为闭包结构
func func09() func() int {
	//1,定义一个局部变量
	a := 0
	//2.定义一个匿名函数,给变量自增并返回
	fun := func() int {
		a++
		return a
	}
	return fun
}

func main() {
	//func01("a")
	//func01("a", "b")
	//func02()
	//func03()
	//func04()
	//func05()
	//func06()
	//func07()
	//func08()
	fmt.Println(func09()())
}
