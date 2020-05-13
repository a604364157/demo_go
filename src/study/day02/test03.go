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

//

func main() {
	func16()
	func17()
	func18()
	func19()
	func20()
}
