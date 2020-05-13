//接口
package main

import "fmt"

//接口的定义
type Phone interface {
	call()
}

type IPhone struct {
}

func (IPhone) call() {
	fmt.Println("iphone call")
}

func func26() {
	var phone Phone
	phone = new(IPhone)
	phone.call()
}

func main() {
	func26()
}
