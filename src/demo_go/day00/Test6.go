package main

import (
	"fmt"
	"math"
)

//常量
const a string = "1"
const b = "1"
const (
	c = 1
	d = 2
)

//批量申明的省略写法，值和前面申明的为一样的时候可以完全简写
const (
	e = 1
	f
	g = 2
	h
)

//iota常量生成器
//常量声明可以使用 iota 常量生成器初始化，
// 它用于生成一组以相似规则初始化的常量，
// 但是不用每行都写一遍初始化表达式。
// 在一个 const 声明语句中，
// 在第一个声明的常量所在的行，
// iota 将会被置为 0，
// 然后在每一个有常量声明的行加一
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//无类型常量
//常量math.Pi无具体类型，而是根据使用场景
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
