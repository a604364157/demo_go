package main

import (
	"demo_go/entity"
	"fmt"
	"time"
)

func main() {
	user := entity.User.C
	user.Say()
	fmt.Println(time.Now().String())
	for true {
		t := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(t)
		time.Sleep(time.Second)
	}
}
