//变量的使用
package main

//变量的声明
func func01() {
	//变量定义
	//var name = value
	//推导式
	//name := value

	var a int
	a = 1
	var b = 2
	c := 3

	//go编译有个特性,未使用的变量会报错
	println(a, b, c)
}

//多变量声明
func func02() {
	var a, b, c int
	a, b, c = 1, 2, 3
	print(a, b, c)

	var e, f, g = 4, 5, 6
	print(e, f, g)

	h, i, j := 7, 8, 9
	print(h, i, j)

	var (
		k int
		l int
		m int
	)
	k, l, m = a, b, c
	print(k, l, m)
}

func main() {
	println("func01----------------------")
	func01()
	println("func02----------------------")
	func02()
}
