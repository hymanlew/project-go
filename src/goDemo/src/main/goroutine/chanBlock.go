package goroutine

import (
	"fmt"
	"time"
)

// 使用 select 可以解决从管道取数据的阻塞问题，即读取未关闭的管道，不会发生死锁异常
func blockNoError() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	//但在实际开发中，可能我们不好确定什么关闭该管道，此时可使用 select 方式可以解决
	check := false
	for {
		if check {
			fmt.Println("都取不到了，退出")
			break
		}
		select {
		//注意: 这里如果 intChan 一直没有关闭，不会一直阻塞而 deadlock，会自动到下一个 case 匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			check = true
		}
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

//goroutine 中使用 recover，解决协程中出现 panic，导致程序崩溃的问题
func catchError() {
	//这里使用 defer + recover，捕获本方法抛出的 panic
	defer func() {
		if err := recover(); err != any(nil) {
			fmt.Println("catchError() 发生错误", err)
		}
	}()

	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang"
}

func testError() {
	go sayHello()
	go catchError()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
