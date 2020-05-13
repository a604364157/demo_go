//结构体(具有任意类型的数据构成的数据集合)
package main

import "fmt"

//定义
type user struct {
	name string
	age  int
	sex  bool
}

func func16() {
	//初始化(顺序赋值)
	a := user{"王小明", 25, true}
	//初始化(指定赋值)
	b := user{name: "刘小红", age: 24, sex: false}
	//new方式,赋默认初始值(new的方式拿到的是指针)
	c := new(user)
	fmt.Println(a, b, c)
	c.name = "小小明"
	fmt.Println(a, b, *c)
}

//结构体的指针
func func17() {
	a := user{"王小明", 25, true}
	fmt.Println(a)
	fmt.Println(&a)
}

//结构体嵌套
type worker struct {
	user         //工作者包含user的全部属性
	wage float64 //工作者添加一个工资属性
}

//结构体嵌套
func func18() {
	//初始嵌套结构体
	a := user{"王小明", 25, true}
	b := worker{a, 5000}
	c := worker{user{"刘小红", 24, false}, 5000}
	fmt.Println(b)
	fmt.Println(c)
	//修改属性(可以通过嵌套属性也可以直接操作属性)
	b.user.age = 26
	fmt.Println(b)
	c.age = 25
	fmt.Println(c)
}

//字段嵌套结构
type boss struct {
	work    worker //老板有一个工作者的字段
	company string //添加一个公司名
}

func func19() {
	var a boss
	a.work = worker{user{"王小明", 25, true}, 10000}
	a.company = "中国航天"
	fmt.Println(a)
}

//提升字段(对于work来说,user变是提升字段)
func func20() {
	a := worker{wage: 5000}
	a.user = user{"王小明", 25, true}
	fmt.Println(a)
}

//结构体是值类型
func func21() {
	fmt.Println("-------------------func21----------------")
	a := user{"王小明", 25, true}
	b := user{"王小明", 25, true}
	//可见两个值相等，但是指针不相等
	fmt.Println(a == b)
	fmt.Println(&a == &b)

	//如果结果体保护的字段不可比较，那么结构体不可比较
	type image struct {
		data map[string]int
	}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := image{m}
	d := image{m}
	//fmt.Println(c == d) 不可比较，编译报错
	fmt.Println(c, d)
	func23(a)
}

//结构体作为参数
func func23(u user) {
	fmt.Println(u)
}

type Employee struct {
	name     string
	salary   int
	currency string
}

//方法定义（方法和函数类似）
func func24(e Employee) {
	fmt.Println(e)
}

//结构体中的函数继承（在test03中，我们定义了user和worker）
//给user定义一个函数
func (u *user) SayHi() {
	fmt.Println("hello ", u.name)
}

func func25() {
	a := user{"王小明", 25, true}
	b := worker{a, 5000}
	//因为worker包含了user，当user定义的函数，worker就自动继承了该函数
	a.SayHi()
	b.SayHi()
}

//函数的重写(注释掉这个函数，执行效果不一样)
func (w worker) SayHi() {
	fmt.Println("Hi ", w.name)
}

func main() {
	func16()
	func17()
	func18()
	func19()
	func20()
	func21()
	emp := Employee{"王小明", 5000, "$"}
	func24(emp)
	func25()
}
