//控制台的输入和输出
package main

import (
	"bufio"
	"fmt"
	"os"
)

//fmt包输出
func func09() {
	//普通打印
	fmt.Print("123", "\n")
	//换行打印
	fmt.Println("123", "456")
	//格式化打印
	fmt.Printf("%s", "123456")
	/*	%v,原样输出
		%T，打印类型
		%t,bool类型
		%s，字符串
		%f，浮点
		%d，10进制的整数
		%b，2进制的整数
		%o，8进制
		%x，%X，16进制
		%x：0-9，a-f
		%X：0-9，A-F
		%c，打印字符
		%p，打印地址*/
}

//fmt包输入
func func10() {
	var a int
	fmt.Println("请输入一个整数")
	fmt.Scanln(&a)
	fmt.Println(a)
}

//bufio包读取
func func11() {
	fmt.Println("请输入一串文字")
	reader := bufio.NewReader(os.Stdin)
	s1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
	fmt.Println("你刚刚输入的是: " + s1)
}

func main() {
	fmt.Println("func09-------------------")
	func09()
	fmt.Println("func10-------------------")
	func10()
	fmt.Println("func11-------------------")
	func11()
}
