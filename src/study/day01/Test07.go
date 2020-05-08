//集合结构-数组
package main

import (
	"fmt"
	"sort"
)

//数组
func func17() {
	//定义数组会给默认值
	var a [10]int
	fmt.Println(a)
	//定义数组,元素个数不能大于数组长度
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	//定义数组,自动匹配元素个数
	c := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(c)
	//数组的遍历1
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	//数组的遍历2
	for x, y := range c {
		fmt.Println(x, y)
	}
}

//多维数组
func func18() {
	a := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a)
	//遍历1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
	//遍历2
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

//数组内是值类型,而不是引用,所以数组a赋值给一个新的变量b,b元素改变不会影响a
func func19() {
	a := []int{1, 2, 3}
	b := a
	b = append(b, 4)
	fmt.Println(a)
	fmt.Println(b)
}

//数组的排序(冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序)
func func20() {
	arr := []int{4, 3, 5, 1, 2}
	fmt.Println(arr)
	//默认使用的快速排序
	sort.Ints(arr)
	fmt.Println(arr)
}

func main() {
	//func17()
	//func18()
	//func19()
	func20()
}
