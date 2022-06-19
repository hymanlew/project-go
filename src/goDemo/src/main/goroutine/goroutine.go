package goroutine

import (
	"fmt"
	"strconv"
	"time"
)

/*
- Go 主线程（也就是进程，一个程序在系统中最少是要有一个进程的），也可称为线程 / 进程。
- 一个 Go 线程上可以起多个协程，可以这样理解，协程是轻量级的线程（编译器上做了优化）。

- 协程的特点：
  - 有独立的栈空间
  - 共享程序堆空间
  - 调度由用户控制（而进程是由操作系统控制的，只是是由用户触发操作）
  - 协程是轻量级的线程

- 主线程是一个物理线程，是直接作用在 cpu 上，是重量级的，非常耗费 cpu 资源。
- 而协程是从主线程开启的，是轻量级的线程，是逻辑态，对资源消耗相对小。
- Golang 的协程机制是其支持高并发的重要特点，它可以轻松的开启上万个协程。
- 其它编程语言的并发机制一般是基于线程的，开启过多的线程，资源耗费大，这里就突显了 Golang 在并发上的优势了。
*/

func goroutinDemo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, goroutine 协程 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	//在主线程(可以理解成进程)中，开启一个 goroutine 协程
	go goroutinDemo()

	for i := 1; i <= 5; i++ {
		fmt.Println("hello, main thread " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
