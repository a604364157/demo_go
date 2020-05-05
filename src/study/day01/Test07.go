package main

import "fmt"

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

func main() {
	func13()
}
