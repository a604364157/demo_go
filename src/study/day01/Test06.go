//流程结构(if)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//if语句
func func12() {
	MIN := 0
	MAX := 100
	var a, b int
	//生成随机整数
	rand.Seed(time.Now().Unix())
	a = rand.Intn(MAX + 1)
	fmt.Println("开始数字炸弹游戏")
	fmt.Printf("已随机生成了一个%d-%d之间的整数,请玩家猜数字\n", MIN, MAX)
	//循环语句
	for a != b {
		fmt.Scanln(&b)
		if a < b {
			if MAX > b {
				MAX = b
			}
			fmt.Printf("数字范围%d-%d\n", MIN, MAX)
		} else if a > b {
			if MIN < b {
				MIN = b
			}
			fmt.Printf("数字范围%d-%d\n", MIN, MAX)
		} else {
			fmt.Println("恭喜你,猜中啦!")
		}
	}
}

func main() {
	func12()
}
