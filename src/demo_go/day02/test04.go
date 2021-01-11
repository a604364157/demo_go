//接口
package main

import "fmt"

//接口的定义//任意类型都实现了空interface
type Phone interface {
	call()
}

type IPhone struct {
}

func (IPhone) call() {
	fmt.Println("iphone call")
}

func func26() {
	var phone Phone
	phone = new(IPhone)
	phone.call()
}

//我们来测试一下 interface的值,先定义一些结构体
type T04a struct {
	name  string
	age   int
	phone string
}

type T04b struct {
	T04a
	school string
}

type T04c struct {
	T04b
	company string
}

//T04a实现一个SayHi方法,然后T04b,T04c自动继承该方法
func (t T04a) SayHi() {
	fmt.Printf("hello %s\n", t.name)
}

//同上
func (t T04a) AddAge() {
	t.age++
}

//T04c重写SayHi方法
func (t T04c) SayHi() {
	fmt.Printf("Hi %s\n", t.name)
}

//定义一个接口T04Com
//包含SayHi(),AddAge()
//被T04a,b,c三个结构体实现
type T04Com interface {
	SayHi()
	AddAge()
}

//测试接口的值
func func27() {
	a := T04b{T04a{"小小明", 3, "123456"}, "幼儿园"}
	b := T04c{T04b{T04a{"王小明", 27, "123456"}, "毕业"}, "中国航天"}
	c := T04c{T04b{T04a{"刘小红", 26, "123456"}, "毕业"}, "中国航天"}
	//定义接口变量i
	var i T04Com
	//i能存储值a
	i = a
	i.SayHi()
	//i能存储值b
	i = b
	i.SayHi()
	//i能存储值c
	i = c
	i.SayHi()
	//切片类
	x := make([]T04Com, 3)
	x[0], x[1], x[2] = a, b, c
	for _, v := range x {
		v.SayHi()
	}
}

//嵌入接口的运用
//先定义一个测试包
type T04Controller struct {
	state int
}

type T04Do interface {
	Get()
	Post()
}

func (c *T04Controller) Get() {
	fmt.Println("GET")
}

func (c *T04Controller) Post() {
	fmt.Println("POST")
}

func func28() {
	var something T04Do
	something = new(T04Controller)
	var t T04Controller
	t.state = 200
	something.Get()
}

//go语言既是静态语言,又有动态语言的便利性,这体现在接口的实现上
//go语言实现接口的方式,是实现接口中的函数,而不是直接实现接口

//接口的断言,因为所有的函数都实现了空接口,当函数的参数是 interface{}时,
//理论可以接收任意对象,那么要获取到真实的对象类型,需要断言
func func29() {
	fmt.Println("-------------------func29-------------------")
	var i interface{} = new(T04a)
	//不安全断言(这里会直接panic掉)
	//s := i.(T04a)
	//fmt.Println(s)
	//安全断言
	s2, ok := i.(T04a)
	fmt.Println(s2, ok)

	//switch 方式
	switch ins := i.(type) {
	case T04a:
		fmt.Println("T04a。。。", ins)
	case T04b:
		fmt.Println("T04b。。。。", ins)
	case T04c:
		fmt.Println("T04c。。", ins)
	default:
		fmt.Println("default。。", ins)
	}
}

func main() {
	func26()
	func27()
	func28()
	func29()
}
