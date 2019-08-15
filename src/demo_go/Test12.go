package main

import "fmt"

//多维数组
func main1() {
	var arr1 [4][2]int
	arr2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	arr3 := [4][2]int{1: {10, 11}, 3: {30, 31}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//切片
	a := [3]int{1, 2, 3}
	//切片索引含头不含尾
	fmt.Println(a, a[1:2])

	//动态构造切片
	b := make([]int, 2)
	c := make([]int, 2, 10)
	fmt.Println(b, c)
	fmt.Println(len(b), len(c))

	//为切片添加元素
	var d []int
	d = append(d, 1)
	d = append(d, 2, 3, 4)
	//追加切片时，切片需要解包，也就是第二种写法的简写
	d = append(d, []int{5, 6, 7}...)
	//扩容的规律是2的倍数

	//追加到切片开头
	d = append([]int{0}, d...)
	d = append([]int{-3, -2, -1}, d...)
	fmt.Println(d)

	//指定位置
	var e []int
	i := 0
	fmt.Println(e)
	e = append(e[:i], append([]int{0}, e[i:]...)...)
	fmt.Println(e)
	e = append(e[:i], append([]int{1, 2, 3}, e[i:]...)...)
	fmt.Println(e)
}

//切片的复制
func main2() {
	const count = 1000
	srcData := make([]int, count)
	//赋值
	for i := 0; i < count; i++ {
		srcData[i] = i
	}

	//引用切片数据
	refData := srcData
	//构造一个新切片
	copyData := make([]int, count)
	//复制
	copy(copyData, srcData)
	//修改远切片数据
	srcData[0] = 999

	//看结果
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[count-1])

	//复制原始数据从4到6
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}

//从切片删除元素
