package goroutine

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
		time.Sleep(time.Millisecond)
	}
	close(intChan)
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		//在遍历时，如果 channel 没有关闭，当遍历到最后时程序会认为可能有数据断续写入，因此就会等待。但如果程序没有数据写入，
		//则会出现 deadlock 的错误。
		//在遍历时，如果 channel 已经关闭，则可以正常遍历数据，遍历完后，就会退出遍历。
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Millisecond)
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	exitChan <- true
	close(exitChan)

}

func workTest() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeToFile(file *os.File, check chan bool) {
	writer := bufio.NewWriter(file)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		num := rand.Intn(1000)
		writer.WriteString(strconv.Itoa(num))
		writer.WriteString(",")
		time.Sleep(time.Millisecond)
	}
	writer.Flush()
	check <- true
}

func readFromFile(file *os.File, ints *[]int, check chan bool) {
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString(',')
		if err == io.EOF {
			break
		}
		s := strings.TrimRight(data, ",")
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error 读取到的数据为 ：" + data)
		} else {
			*ints = append(*ints, i)
		}
	}
	sort.Ints(*ints)
	check <- true
}

func writeToFile2(file *os.File, ints *[]int) {
	writer := bufio.NewWriter(file)
	for _, v := range *ints {
		writer.WriteString(strconv.Itoa(v))
	}
	writer.Flush()
}

func checkChannel(check chan bool) {
	for {
		v, ok := <-check
		if ok {
			fmt.Printf("检查完成：%v\n", v)
			break
		}
	}
}

func WorkTest2() {
	filPath := "/channel.txt"
	var file *os.File
	if f, err := os.Stat(filPath); f != nil || os.IsExist(err) {
		//os.O_TRUNC 清空已经存在的文件
		file, err = os.OpenFile("/channel.txt", os.O_RDWR|os.O_TRUNC, 0777)
		if err != nil {
			panic(any("打开文件失败 " + err.Error()))
		}
	} else {
		file, err = os.Create(filPath)
		if err != nil {
			panic(any("创建文件失败"))
		}
	}
	defer file.Close()

	//var check chan bool
	check := make(chan bool, 1)
	go writeToFile(file, check)
	checkChannel(check)

	ints := make([]int, 0, 1000)
	file, _ = os.Open(filPath)
	go readFromFile(file, &ints, check)
	checkChannel(check)

	filPath = "/channel2.txt"
	file2, err2 := os.Create(filPath)
	if err2 != nil {
		panic(any("创建写入文件失败"))
	}
	defer file2.Close()
	go writeToFile2(file2, &ints)

	for _, v := range ints {
		fmt.Println(v)
	}
}
