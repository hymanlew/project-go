package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 需求：计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。最后显示出来。要求使用goroutine完成
var (
	myMap = make(map[int]int, 10)

	//lock 是一个被声明的全局的互斥锁
	//sync 是 synchornized 同步包，Mutex 是互斥。但是此包只适用于低效率的线程，高效率的线程通信还是要用 channel 的
	lock sync.Mutex
)

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func calTest(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 不加锁时，直接运行会产生 concurrent map writes 异常，即存在资源竞争问题。如果要不抛出异常，而要输出计算结果。可以使用命令运行程序，
	// go build -race main.go
	//myMap[n] = res

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func DoTest() {
	// 这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go calTest(i)
	}

	//休眠 10 秒钟，等待协程处理完成
	time.Sleep(time.Second * 5)

	//这里加入锁，是因为在实际中不可能硬编码让主线程等待。因此实际中主线程并不知道所有协程是否执行结束，仍有可能出现资源争夺问题，
	//因此加入互斥锁即可解决问题。
	//但通过对全局变量加锁同步来实现通讯，也并不利于多个协程对全局变量的读写操作。因此还是要使用 channel
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
