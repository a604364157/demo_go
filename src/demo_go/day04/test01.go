//并发编程
/*并发编程，什么是并行，并发，串行。Go语言如何实现并发编程
goroutine的使用。runtime包、sync包的介绍。channel通道的使用，
以及缓冲通道，定向通道。select语句，time包中和并发编程相关的函数介绍*/
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//go的并发是基于协程 goroutine
func func01() {
	//准备一个最简单的并发测试
	p1 := func() {
		for true {
			fmt.Println("1111111111")
		}
	}
	p2 := func() {
		for true {
			fmt.Println("2222222222")
		}
	}
	go p1()
	go p2()
	time.Sleep(10 * time.Second)
}

//runtime 调度器
/*
	NumCPU：返回当前系统的 CPU 核数量
	GOMAXPROCS：设置最大的可同时使用的 CPU 核数
	Gosched：让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
    Goexit：退出当前 goroutine(但是defer语句会照常执行)
	NumGoroutine：返回正在执行和排队的任务总数
	GOOS：目标操作系统
	runtime.GC:会让运行时系统进行一次强制性的垃圾收集
	GOROOT :获取goroot目录
*/
func func02() {
	fmt.Println("GOROOT-->", runtime.GOROOT())
	fmt.Println("操作系统:", runtime.GOOS)
	fmt.Println("逻辑CPU核心数", runtime.NumCPU())
}

//Gosched：让当前线程让出 cpu 以让其它线程运行
func func03() {
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("1111111111")
		}
	}()
	//偶然发现还是会有2在1之前输出的,还是CPU的切换在线程数
	for i := 0; i < 6; i++ {
		runtime.Gosched()
		fmt.Println("2222222222")
	}
}

//Goexit：退出当前 goroutine(但是defer语句会照常执行)
func func04() {
	test := func() {
		defer fmt.Println("test defer----------")
		//终止此函数
		runtime.Goexit()
		fmt.Println("test end----------")
	}
	//可以看到,运行到runtime.Goexit()后,除了已加载的defer,其他均不会再执行
	go func() {
		fmt.Println("开始------------")
		test()
		fmt.Println("结束------------")
	}()

	time.Sleep(5 * time.Second)
}

//临界资源
func func05() {
	a := 1
	go func() {
		a = 2
		fmt.Println("sub      ", a)
	}()
	a = 3
	time.Sleep(1)
	fmt.Println("func05      ", a)
	//可见两个输出都是2,出现数据异常
}

//我们来模拟一个卖票的场景
func func06() {
	num := 50
	sale := func(name string) {
		for {
			if num > 0 {
				// 睡眠随机数
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				num--
				fmt.Println(name, "售出,余票:", num)
			} else {
				fmt.Println(name, "没有票了...")
				break
			}
		}
	}
	go sale("售票员1")
	go sale("售票员2")
	go sale("售票员3")
	go sale("售票员4")

	time.Sleep(10 * time.Second)
	//这里可以看到把票卖成负数, 是一个协程运行到num>0后,我们延迟了,这时其他协程改变了num的值
}

//我用使用锁再实现上面的例子
func func07() {
	//创建锁头
	var wg sync.WaitGroup
	var matex sync.Mutex
	num := 20
	sale := func(name string) {
		defer wg.Done()
		for {
			matex.Lock()
			if num > 0 {
				// 睡眠随机数
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				num--
				fmt.Println(name, "售出,余票:", num)
			} else {
				matex.Unlock()
				fmt.Println(name, "没有票了...")
				break
			}
			matex.Unlock()
		}
	}
	wg.Add(4)
	go sale("售票员1")
	go sale("售票员2")
	go sale("售票员3")
	go sale("售票员4")
	wg.Wait()
	//我们通过锁,锁住了逻辑代码块,代码块在同一时间只能有一个协程执行
}

//说到了锁,我们再来说一下WaitGroup同步等待组
func func08() {
	//Add(),Done(),Wait()函数解释
	//Add(),等待组中执行协程得数量
	//Done(),某个协程执行完毕后,让WaitGroup的counter数值减1 本质是 Add(-1)
	//Wait(),让当前协程阻塞,直到等待组内协程执行完毕
	//使用实例参考func07
}

//func07还使用到了 sync.Mutex 互斥锁
func func09() {
	//Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。
	//Mutex 是最简单的一种锁类型，互斥锁
	//Lock锁定当前的共享资源，Unlock进行解锁
	//使用实例参考func07
}

//ync.RWMutex 读写锁(只能有一个协程获得写锁,可以有多个协程获取读锁)
func func10() {
	rwMutex := new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	r := func(i int) {
		defer wg.Done()
		fmt.Println(i, "read start")
		rwMutex.RLock() //加上读锁
		fmt.Println(i, "reading")
		time.Sleep(3 * time.Second)
		rwMutex.RUnlock() //解除读锁
		fmt.Println(i, "read over")
	}
	fmt.Println(r)
	w := func(i int) {
		defer wg.Done()
		fmt.Println(i, "write start")
		rwMutex.Lock()
		fmt.Println(i, "writing")
		time.Sleep(3 * time.Second)
		rwMutex.Unlock()
		fmt.Println(i, "write over")
	}
	fmt.Println(w)

	wg.Add(3)

	//模拟多个读操作
	//go r(1)
	//go r(2)
	//go r(3)

	//模拟多个写操作
	go w(1)
	go r(2)
	go w(3)

	wg.Wait()
	fmt.Println("func10 执行完毕")

	/*读锁不能阻塞读锁
	读锁需要阻塞写锁，直到所有读锁都释放
	写锁需要阻塞读锁，直到所有写锁都释放
	写锁需要阻塞写锁*/
}

func main() {
	//func01()
	//func02()
	//func03()
	//func04()
	//func05()
	//func06()
	//func07()
	func10()
}
