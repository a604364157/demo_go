//集合结构-Map
package main

import "fmt"

// map的定义
func func24() {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("a=", a)
	m := make(map[string]int)
	fmt.Println("m=", m)
	//添加元素
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println("m=", m)
	//长度
	fmt.Println("m len=", len(m))
	//遍历输出
	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}
	//检测元素是否存在
	_, ok := m["d"]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	//删除元素
	delete(m, "c")
	fmt.Println("m=", m)
	//map是引用类型,新的变量改变会反应到原对象上
	m2 := m
	m2["c"] = 3
	fmt.Println("m=", m)
	fmt.Println("m2=", m2)
}

func main() {
	func24()
}
