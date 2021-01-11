package main

import (
	"fmt"
	"strconv"
)

//go语言字符串和数字的转换

func main() {
	a := 123
	b := fmt.Sprintf("%d", a)
	//两种转换方式转字符串
	fmt.Println(b, strconv.Itoa(a))

	//按进位制格式化数字
	fmt.Println(strconv.FormatInt(int64(a), 2))
	c := fmt.Sprintf("a = %b", a)
	fmt.Println(c)

	//两种方式将字符串转数字
	d, err := strconv.Atoi("123")
	e, err := strconv.ParseInt("123", 10, 64) //十进制，最长64位
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d, e)
}
