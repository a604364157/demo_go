//流程结构
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

//switch
func func13() {
	var a int
	fmt.Println("请输入你的语文成绩")
	fmt.Scanln(&a)
	//注:go的switch不需要break来跳出,相反,如果需要贯通,需要fallthrough关键字
	switch a {
	case 90:
		fmt.Println("你的成绩为: 优秀")
	case 80:
		fmt.Println("你的成绩为: 良好")
	case 60, 70:
		fmt.Println("你的成绩为: 及格")
	default:
		fmt.Println("你的成绩为: 不及格")
	}
}

//type-switch
func func14(x interface{}) {
	switch t := x.(type) {
	case nil:
		fmt.Printf("X的类型: %T", t)
	case string:
		fmt.Printf("X的类型: %s", "string")
	case bool:
		fmt.Printf("X的类型: %s", "bool")
	case int:
		fmt.Printf("X的类型: %s", "int")
	case uint:
		fmt.Printf("X的类型: %s", "uint")
	case float32, float64:
		fmt.Printf("X的类型: %s", "float")
	default:
		fmt.Printf("X的类型: %s", "非基本类型")
	}
}

//循环语句(for)
func func15() {
	//基本循环
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	//range迭代(range可以迭代:切片,列表, 元组, 字典等集合结构)
	a := "0123456789"
	for i := range a {
		fmt.Print(i)
	}
	fmt.Println()
	//break跳出循环
	for i := range a {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}
	fmt.Println()
	//continue进入下一次循环
	for i := range a {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
}

//goto语句
func func16() {
	a := 10

LOOP:
	for a < 20 {
		if a == 15 {
			a++
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}

func main() {
	//func12()
	//func13()
	//func14('1')
	//func15()
	func16()
}
