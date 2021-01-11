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
//错误是 一个带有 Error() string 函数的接口
//怎么从错误中提取更多信息,对比func31
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

//自定义错误(这里只是一个很简单的例子)
type T06Error struct {
	err  string
	code string
}

func (e *T06Error) Error() string {
	return fmt.Sprintf("error code %s; msg: %s\n", e.code, e.err)
}

func func33(i int) (int, error) {
	if i < 0 {
		return 0, &T06Error{"传入参数小于0了", "001"}
	}
	return i, nil
}

func func34() {
	i, err := func33(-1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(i)
}

func main() {
	//func30()
	func31()
	func32("golangbot123.com")
	func32("www.baidu.com")
	func34()
}
