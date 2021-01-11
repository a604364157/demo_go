package main

import (
	"fmt"
	"reflect"
)

//type关键字，类型定义

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

//在结构体成员嵌入时使用别名
//定义商标
type Brand struct {
}

//给它添加show函数
func (t Brand) show() {

}

//定义一个别名
type FakeBrand = Brand

//定义一个类型（车辆）
type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)

	var b IntAlias
	fmt.Printf("b type: %T\n", b)

	//运行结果
	//a type: main.NewInt
	//b type: int

	var v Vehicle
	v.FakeBrand.show()
	ta := reflect.TypeOf(v)
	for i := 0; i < ta.NumField(); i++ {
		//v的成员信息
		f := ta.Field(i)
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}
