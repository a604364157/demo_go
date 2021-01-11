//type关键字
package main

//定义结构体
type T05a struct {
	name string
}

//定义接口
type T05Inter interface {
	test()
}

//定义其他类型 type 类型名 Type

//定义函数
type T05Func func(int, int) string

//定义类型别名   type 别名 = Type
type byte = uint8
type rune = int32

func main() {

}
