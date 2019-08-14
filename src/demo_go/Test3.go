package main

import (
	"fmt"
	"math"
)

func main() {
	//输出各数值范围
	fmt.Println("int8 range: ", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range: ", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range: ", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range: ", math.MinInt64, math.MaxInt64)

	//初始一个32位整型值
	var a int32 = 1047483647
	//将a转为16进制，发生数值截断
	b := int16(a)
	//输出变量的16进制和10进制形式
	fmt.Printf("int16: 0x%x   %d\n", b, b)
	//整型截断是一种很隐形的BUG，需要特别注意

	//将常量保存为float32类型(这个常量是圆周率)
	var c float32 = math.Pi
	//转换为int类型，发生精度丢失
	fmt.Println(int(c))
}
