package goroutine

import "fmt"

func chanDemo() {
	//- channel 本质上就是一个 FIFO 队列，且多 goroutine 访问时不需要加锁，即 channel 本身就是线程安全的。
	//- channel 管道是有类型的，只能存放指定的数据类型，如一个 string 的 channel 只能存放 string 类型数据。
	//- channel 管道是引用类型，且必须初始化后才能写入数据，即 make 后才能使用。

	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n", intChan, &intChan)

	//2. 向管道写入数据时，不能超过其容量。且数据放满后，就不能再放入了。
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	//3. 如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	//num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)
}

type Cat struct {
	Name string
	Age  int
}

func ChanType() {
	//定义一个存放任意数据类型的管道 3个数据
	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	//我们希望得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)

	//当 channel 中存入的是空接口类型时，取出数据使用时，要注意做类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}

func forRange() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)

	//使用内置函数 close 可以关闭 channel，且该通道必须为双向或只发送的。它应当只由发送者执行，而不应由接收者执行，其是在最后发送的
	//值被取出后关闭该通道。当 channel 关闭后，就不能再向 channel 写数据了，但仍然可以从该 channel 读取数据。

	//通道关闭后，不能够再写入数到channel
	//intChan<- 300

	//当管道关闭后，读取数据是可以的
	x, ok := <-intChan
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("error")
	}

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//遍历管道不能使用普通的 for 循环，因为每执行一次，通道的长度就减少一次
	// for i := 0; i < len(intChan2); i++ {
	// }

	//在遍历时，如果 channel 没有关闭，当遍历到最后时程序会认为可能有数据断续写入，因此就会等待。但如果程序没有数据写入，则会出现 deadlock 的错误
	//在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}

func chanRead() {
	//管道可以声明为只读或者只写。默认管道是双向的，即可读可写
	//var chan1 chan int

	// 声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//num := <-chan2 //error
	fmt.Println("chan2=", chan2)

	// 声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	//chan3<- 30 //err
	fmt.Println("num2", num2)

}
