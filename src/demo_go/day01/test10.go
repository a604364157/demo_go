//字符串
package main

import (
	"fmt"
	"strings"
)

func func25() {
	//定义
	a := "123456789"
	//单个字符
	for _, v := range a {
		fmt.Printf("%d,", v)
		fmt.Printf("%c,", v)
	}
	//strings包,包含常见字符串的操作函数
	//判断前缀
	fmt.Println(strings.HasPrefix(a, "1"))
	//判断后缀
	fmt.Println(strings.HasSuffix(a, "9"))
	//字符索引
	fmt.Println(strings.Index(a, "2"))
	fmt.Println(strings.LastIndex(a, "2"))
	//包含
	fmt.Println(strings.Contains(a, "1"))
	//替换(n为替换次数,-1则全部替换)
	fmt.Println(strings.Replace(a, "1", "0", -1))
	fmt.Println(strings.ReplaceAll(a, "1", "0"))
	//非重复字符的次数
	fmt.Println(strings.Count(a, "0"))
	//重复字符串(count为重复写次数)
	fmt.Println(strings.Repeat(a, 2))
	//大小写切换
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	//去空格
	fmt.Println(strings.TrimSpace(a))
	//去指定字符
	fmt.Println(strings.Trim(a, "1"))
}

func main() {
	func25()
}
