//集合结构-切片
package main

import "fmt"

//切片的定义
func func21() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("s=", s)
	//s := arr[startIndex:endIndex] 索引的一向规则,含头不含尾
	s1 := s[4:]
	fmt.Println("s1=", s1)
	s2 := s[:4]
	fmt.Println("s2=", s2)
	s3 := s[3:4]
	fmt.Println("s3=", s3)
	//切片的修改(切片是数组的映射,数据修改会反应在原数组上)
	s2[0] = 0
	fmt.Println("s=", s)
	fmt.Println("s2=", s2)
}

//切片的长度合容量
func func22() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s len=%d, cap=%d s=%v\n", len(s), cap(s), s)
	s1 := s[4:]
	fmt.Printf("s1 len=%d, cap=%d s=%v\n", len(s1), cap(s1), s1)
	s2 := make([]int, 3, 5)
	fmt.Printf("s2 len=%d, cap=%d s=%v\n", len(s2), cap(s2), s2)
	var s3 []int
	fmt.Printf("s3 len=%d, cap=%d s=%v\n", len(s3), cap(s3), s3)
}

//append() 和 copy() 函数
func func23() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s=", s)
	s = append(s, 6)
	fmt.Println("s=", s)
	s = append(s, 7, 8, 9)
	fmt.Println("s=", s)
	s2 := make([]int, len(s), cap(s)*2)
	copy(s2, s)
	fmt.Printf("s2 len=%d, cap=%d s=%v\n", len(s2), cap(s2), s2)
}

func main() {
	//func21()
	//func22()
	func23()
}
