package main

import (
	"fmt"
	"math/rand"
	"time"
)

//例1：生产者每秒生成一个数据，并通过通道传递给消费者
//			生产者使用两个goroutine并发，消费者在main函数的goroutine中进行处理

//生产者
func producer(header string, channel chan<- string) {
	for true {
		channel <- fmt.Sprintf("%s：%v", header, rand.Int31())
		//协程等待1秒
		time.Sleep(time.Second)
	}
}

//消费者
func customer(channel <-chan string) {
	for true {
		msg := <-channel
		fmt.Println(msg)
	}
}

func main() {
	//创建一个通道
	channel := make(chan string)
	//创建生产者并且 并发goroutine
	go producer("A", channel)
	go producer("B", channel)
	//消费
	customer(channel)
}
