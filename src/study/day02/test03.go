//指针
package main

import "fmt"

//指针声明:*T是指针变量类型
func func20() {
	var a int = 20
	var b *int
	b = &a
	fmt.Println("a的地址是", b)
	//可见*T 指针和值的切换
	fmt.Println("a的值是", *b)
}

type name int

type test struct {
	a int
	b bool
	name
}

//指针的实际运用
func func21() {
	a := test{1, false, 2}
	b := &a
	fmt.Println(a.a, a.b, a.name, &a, b.a, &b, (*b).a)
}

//空指针
func func22() {
	//go的空表示为 nil
}

// 指针的操作
func func23() {
	//获取指针的值(语法 *T)
	a := 123
	b := &a
	fmt.Println("b=", b)
	fmt.Println("*b=", *b)

	//通过指针改变变量的值
	*b++
	fmt.Println("a=", a)

	//通过指针传递参数
	fun := func(a *int) *int {
		*a++
		return a
	}
	fmt.Println(*fun(b))

	//切片的指针(以下写法是示例,一般这样的逻辑使用切片实现)
	fun2 := func(arr *[5]int) {
		(*arr)[0] = 2
	}
	c := [5]int{1, 2, 3, 4, 5}
	fun2(&c)
	fmt.Println("c=", c)
}

//指针的指针(**T)
func func24() {
	var a int
	var b *int
	var c **int
	a = 1
	b = &a
	c = &b
	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Println("a=", a, "*b=", *b, "**c=", **c)
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

//指针作为函数的参数
func func25() {
	a := 100
	b := 200
	swap(&a, &b)
	fmt.Println("a,b=", a, b)
}

func main() {
	func20()
	func21()
	func23()
	func24()
	func25()
}
