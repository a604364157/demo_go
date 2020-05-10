//基本数据类型
package main

/*1、整数型

int8
有符号 8 位整型 (-128 到 127)
长度：8bit

int16
有符号 16 位整型 (-32768 到 32767)

int32
有符号 32 位整型 (-2147483648 到 2147483647)

int64
有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

uint8
无符号 8 位整型 (0 到 255)
8位都用于表示数值：

uint16
无符号 16 位整型 (0 到 65535)

uint32
无符号 32 位整型 (0 到 4294967295)

uint64
无符号 64 位整型 (0 到 18446744073709551615)*/

/*2、浮点型

float32

IEEE-754 32位浮点型数

float64

IEEE-754 64位浮点型数

complex64

32 位实数和虚数

complex128

64 位实数和虚数*/

/*3、其他

byte

类似 uint8

rune

类似 int32

uint

32 或 64 位

int

与 uint 一样大小

uintptr

无符号整型，用于存放一个指针*/

//bool型

//string型

func func07() {
	var a bool = true
	println(a)
	var c int8 = 127
	println(c)
	var d int16 = 32767
	println(d)
	var e int32 = 2147483647
	println(e)
	var f int64 = 9223372036854775807
	println(f)
	//int是跟随底层系统的(32位或者64位)
	var b int = int(f)
	println(b)

	var h uint8 = 255
	println(h)
	var i uint16 = 65535
	println(i)
	var j uint32 = 4294967295
	println(j)
	var k uint64 = 18446744073709551615
	println(k)
	//uint是跟随底层系统的(32位或者64位)
	var l uint = uint(k)
	println(l)

	var m float32 = 3.14
	println(m)
	var n float64 = 3.14
	println(n)

	var o complex64 = 1
	println(o)
	var p complex128 = 1
	println(p)

	var q string = "123456789"
	println(q)
}

func main() {
	println("func07----------------------")
	func07()
}
