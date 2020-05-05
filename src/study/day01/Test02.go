//常量的使用
package main

import "fmt"

//常量的声明(ps:go是以包为运行单位,包内的函数名不能重复)
func func03() {
	//显式
	const A string = "A"
	//隐式
	const B = "B"

	println(A, B)
}

//例
func func04() {
	const HEIGHT = 20
	const WIDTH = 20
	area := HEIGHT * WIDTH
	fmt.Printf("面积为: %d", area)
}

//模拟枚举
func func05() {
	const (
		ONE   = 1
		TWO   = 2
		THREE = 3
	)
	println(ONE, TWO, THREE)
}

func func06() {
	//iota 自增
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func main() {
	//这里发现一个有趣的现象,默认输出println和fmt.Println是两个独立的线程
	println("func03----------------------")
	func03()
	println("func04----------------------")
	func04()
	println("func05----------------------")
	func05()
	println("func06----------------------")
	func06()
}
