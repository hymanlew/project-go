package goroutine

import (
	"fmt"
	"runtime"
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

MPG 原理图串灰色的队列的 goroutine 协程并没有运行，而是处于 ready 的就绪态，等待被 P 调度。P 维护着这个队列，称为 runqueue。
在 go中使用 go function 就可能启动一个协程，即在 runqueue 队列的末尾追加一个 goroutine。而 P 就在下一调度点从 runqueue 中取出一个
goroutine 执行。

为何要维护多个上下文 P（即维护多个线程），因为当一个 OS 线程被阻塞时，P 可以转而投奔另一个 OS 线程。即调度器需要保证有足够的线程来运行
所有的 context P。

另一种情况是某个线程 P 所分配的任务 G 很快就执行完了（分配不均），这就导致了一个上下文 P 闲着没事而系统却依然忙碌。但如果 global runqueue
也没有任务 G 了，则 P 就不得不从其他上下文 P 那里拿一些 G 来执行。且一般来说，如果上下文 P 从其他上下文 P 那里偷任务的话，会偷 runqueue
的一半，这就确保了每个 OS 线程都能被充分利用。
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

/*
为了充分了利用多 cpu 的优势，在 Golang 程序中，通常需要设置运行的 cpu 数目。
- 在 go 1.8 后，底层默认是让程序运行在多个核上，可以不用设置了。
- 在 go 1.8 前，还是要设置一下，可以更高效的利用 cpu。
*/
func cpuTest() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
