package main

import (
	"fmt"
	"sort"
	"sync"
)

//map的概念

func main4() {
	var m1 map[string]int
	fmt.Println(m1)
	var m2 map[string]int
	m1 = map[string]int{"a": 1, "b": 2}
	m2 = m1
	fmt.Println(m2)

	m3 := make(map[string]float32)
	m3["a"] = 1.2
	m3["b"] = 1.3
	fmt.Println(m3)
	m2["c"] = 3
	fmt.Println(m2)
	fmt.Println(m1)

	m4 := make(map[int][]int)
	m5 := make(map[int]*[]int)
	fmt.Println(m4)
	fmt.Println(m5)
}

//map的遍历
func main5() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}

	scenes := make(map[string]int)
	// 准备map数据
	scenes["route"] = 66
	scenes["brazil"] = 4
	scenes["china"] = 960
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scenes {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)
}

//map删除元素和清空
func main6() {
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//go语言没有提供清空map的函数
	//现在的做法是直接构建一个新的map
}

//sync.Map
func main() {
	tm1()
}

//测试普通map并发问题
//fatal error: concurrent map read and map write
func tm1() {
	m := make(map[int]int)
	//开启一个并发
	go func() {
		for true {
			m[1] = 1
		}
	}()

	//开启第二个并发
	go func() {
		for true {
			_ = m[1]
		}
	}()

	//开启一个死循环，让程序不会停止
	i := 0
	for true {
		i++
	}
}

//并发安全的map
func tm2() {
	//声明
	var scene sync.Map
	//添加数据
	scene.Store("a", 1)
	scene.Store("b", 2)
	scene.Store("c", 3)
	//获取数据
	fmt.Println(scene.Load("b"))
	//删除数据
	scene.Delete("b")
	//遍历
	scene.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
