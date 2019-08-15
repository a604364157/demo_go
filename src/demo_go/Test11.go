package main

import "fmt"

//数组

func main() {
	//定义三个整数的数组
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//定义带初始值的
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	fmt.Println(b)
	fmt.Println(c[2])

	//使用...代替初始长度，根据长度自动定义
	d := [...]int{1, 2, 3}
	fmt.Println(d)

	//比较两个数组
	e := [2]int{1, 2}
	f := [...]int{1, 2}
	g := [2]int{1, 3}
	fmt.Println(e == f, e == g, f == g)
	h := [3]int{1, 2}
	fmt.Println(h)
	//fmt.Println(e == h)//编译错误，数组长度不一致

	//遍历数组
	tt := [3]string{"a", "b", "c"}
	for i, v := range tt {
		fmt.Println(i, v)
	}
}
