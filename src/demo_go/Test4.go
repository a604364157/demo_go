package main

import (
	"flag"
	"fmt"
)

//指针的用法
func main() {
	cat := 1
	str := "banana"
	fmt.Printf("%p %p", &cat, &str)
	//%p就是输出指针，&后面跟变量，这个与c语言的写法一致，代表变量地址
	fmt.Println()

	//准备一个字符串
	house := "malibu point 10880, 90265"
	//取字符串地址
	ptr := &house
	//打印变量类型
	fmt.Printf("ptr type: %T\n", ptr)
	//打印变量指针地址
	fmt.Printf("ptr addr: %p\n", ptr)

	//对指针取值(与&相反的是，&取地址，*根据地址取值)
	value := *ptr
	//取值后值类型
	fmt.Printf("value type: %T\n", value)
	//打印变量值
	fmt.Printf("value: %s\n", value)

	//测试指针交换值
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	//测试指针获取命令行信息
	parse()

	//测试指针new*()的创建方式
	testNew()
}

//使用指针修改值

//交换函数
func swap(a, b *int) {
	//取a指针的值，赋给临时变量t
	t := *a
	//取b指针的值，赋给a指针的变量
	*a = *b
	//将t(也就是a指针的值)，赋给b指针的变量
	*b = t
}

//使用指针获取命令行输入信息
func parse() {
	mode := flag.String("mode", "", "process mode")
	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*mode)
}

//创建指针的另一种方法  new()函数
func testNew() {
	str := new(string)
	*str = "ninja"
	fmt.Println(*str)
}
