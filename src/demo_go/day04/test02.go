//通道channel 不要通过共享内存来通信，而应该通过通信来共享内存
//通道是什么，通道就是goroutine之间的通道。它可以让goroutine之间相互通信
package main

import (
	"fmt"
	"strconv"
	"time"
)

func func11() {
	//通道得声明
	a := make(chan int)
	fmt.Println(a)
	//channel的数据类型
	fmt.Printf("%T,%p\n", a, a)
	//通道的语法使用(使用通道必须是一个协程写入,存在其他协程读取,否则会死锁,而且写入和读取默认是阻塞的)
	go func() {
		//写数据
		a <- 1
	}()
	//读数据
	v, ok := <-a
	fmt.Println(v, ok)
}

//通道的使用(阻塞示例)
func func12() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine正在执行, i=", i)
			time.Sleep(300 * time.Millisecond)
		}
		ch <- true
	}()
	data := <-ch
	//可以发现,当f12读取通道时,被阻塞,等待子协程执行完毕,写入通道后才会执行
	fmt.Println(data)
}

//通道使用例子
func func13() {
	//计算数据的数字和
	s := func(num int, ch chan int) {
		sum := 0
		for num != 0 {
			d := num % 10
			sum += d
			num /= 10
		}
		ch <- sum
	}

	//计算数据的数字平方和
	c := func(num int, ch chan int) {
		sum := 0
		for num != 0 {
			d := num % 10
			sum += d * d
			num /= 10
		}
		ch <- sum
	}

	num := 123

	ch1 := make(chan int)
	ch2 := make(chan int)

	go s(num, ch1)
	go c(num, ch2)

	data1, data2 := <-ch1, <-ch2
	fmt.Println(data1, data2)
}

//关闭通道
func func14() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		time.Sleep(300 * time.Millisecond)
		//判断通过是否关闭
		v, ok := <-ch
		if !ok {
			fmt.Println("读取完成")
			break
		}
		fmt.Println("正在读取数据,当前数据为:", v)
	}
}

//通道的范围循环
func func15() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()
	// for循环的for range形式可用于从通道接收值，直到它关闭为止
	for v := range ch {
		fmt.Println("读取到数据:", v)
	}
}

//缓冲通道
func func16() {
	//非缓冲
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	//缓冲
	ch2 := make(chan int, 10)
	fmt.Println(len(ch2), cap(ch2))

	ch3 := make(chan string, 5)
	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- "数据:" + strconv.Itoa(i)
			fmt.Println("子goroutine,写入了数据", i)
		}
		close(ch3)
	}()
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch3
		if !ok {
			fmt.Println("数据读取完毕")
			break
		}
		fmt.Println("读取到", v)
	}
	//可见,带缓冲的通道,待缓冲满了,只能被读取后才能继续写
}

//定向通道
func func17() {
	//既可以读又可以写的通道称为双向通道

	//这里我们学习一下单项通道
	w := func(data string, ch chan<- string) {
		fmt.Println("已写入", data)
		ch <- data
	}

	r := func(ch <-chan string) {
		data := <-ch
		fmt.Println("读取到", data)
	}

	ch := make(chan string)

	go w("123456789", ch)
	go r(ch)
	time.Sleep(1 * time.Second)
}

//time包中的通道使用(定时器Timer)
func func18() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T\n", timer)

	fmt.Println(time.Now())
	//此处在等待channel中的信号，执行此段代码时会阻塞3秒
	ch := timer.C
	fmt.Println(<-ch)
}

//timer.Stop
func func19() {
	timer := time.NewTimer(5 * time.Second)

	//开一个协程
	go func() {
		<-timer.C
		fmt.Println("timer 结束")
	}()

	//
	time.Sleep(3 * time.Second)
	stop := timer.Stop()
	if stop {
		fmt.Println("timer 停止")
	}
}

//time.After
func func20() {
	ch := time.After(5 * time.Second)
	fmt.Println(time.Now())
	t := <-ch
	fmt.Println(t)
}

// select语句
//select 是 Go 中的一个控制结构。select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行
func func21() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	select {
	case num := <-ch1:
		fmt.Println("ch1中取到数据:", num)
	case num2 := <-ch2:
		fmt.Println("ch2中取到数据:", num2)
	}
}

//select语句结合timer实现超时操作
func func22() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//这里因为一直没有对管道写入,导致 1, 2一直不能执行,最终走向超时
	select {
	case <-ch1:
		fmt.Println("case1可以执行")
	case <-ch2:
		fmt.Println("case2可以执行")
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	}
}

func main() {
	//func11()
	//func12()
	//func13()
	//func14()
	//func15()
	//func16()
	//func17()
	//func18()
	//func19()
	//func20()
	//func21()
	func22()
}
