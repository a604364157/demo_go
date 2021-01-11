package main

import "fmt"

//模拟枚举
//现在go还没有枚举，一般根据常量模拟

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

//用iota生成非+1的常量
const (
	//这里进行移位，每次往左移1位
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

//将枚举值转换为字符串
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	var weapon = Blower
	fmt.Println(weapon)

	fmt.Printf("%d %d %d %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)

	fmt.Printf("%s %d", CPU, CPU)
}
