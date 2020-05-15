//错误
package main

import (
	"fmt"
	"net"
	"os"
)

//错误演示
func func30() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Name())
}

//错误捕捉
func func31() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("file", err.Path, "not found")
		return
	}
	fmt.Println(f.Name())
}

//分类捕捉
func func32(s string) {
	addr, err := net.LookupHost(s)
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			//无效的域名
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func main() {
	//func30()
	func31()
	func32("golangbot123.com")
	func32("www.baidu.com")
}
