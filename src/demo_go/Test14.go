package main

import (
	"container/list"
	"fmt"
)

//list列表

func main() {
	l := list.New()

	//尾部添加元素
	l.PushBack("z")
	//头部添加(保存一个句柄)
	element := l.PushFront("b")
	//指定元素之后
	l.InsertAfter("c", element)
	//在指定元素之前
	l.InsertBefore("a", element)
	fmt.Println(l.Len())
	//遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
