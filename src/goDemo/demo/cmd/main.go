package main

import (
	"fmt"
	"strconv"
)

type A2 struct {
	Name string
	Age  int
}

func (a *A2) String() string {
	i := a.Name + " = " + strconv.Itoa(a.Age)
	return i
}

/*
  每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，被 Go 运行框架调用。因此通常可以在 init 函数中完成初始化工作。
  如果一个文件同时包含全局变量，init 函数和 main 函数，则执行流程是：全局变量定义 --> init 函数 --> main 函数。
  如果源文件的引入包中，也有全局变量定义或 init 函数时，则就会先执行引入包中的变量及 init 函数。

  单元测试相关：
  https://studygolang.com/articles/4104
  go get -v github.com/uudashr/gopkgs/cmd/gopkgs
  go get -v github.com/mdempsky/gocode

  github.com/mdempsky/gocode (download)
  github.com/mdempsky/gocode/gbimporter
  github.com/mdempsky/gocode/lookdot
  github.com/mdempsky/gocode/suggest
  github.com/mdempsky/gocode

  go get -v github.com/sqs/goreturns //能自动引入包
*/
func main() {

	fmt.Println()
	a := A2{
		"tom",
		10,
	}
	fmt.Println(&a)

	//fmt.Println()
	//inoutput.Scan()

	//fmt.Println()
	//funct.Test5()

	//fmt.Println()
	//funct.Test7()

	fmt.Println("主程序结束。。。。")
}
