package main

import (
	"fmt"
	"goDemo/src/main/funct"
	"strconv"
)

type A2 struct {
	Name string
	Age int
}

func (a *A2) String() string {
	i := a.Name + " = " + strconv.Itoa(a.Age)
	return i
}

func main()  {
	// 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用。
	// 可简单理解成，首字母大写是公开的，小写是私有的，在 golang 中没有 public，private 等关键字。
	//datatype.Data()

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
	//fmt.Println(utils.Lengh("hello 你好"))

	//fmt.Println()
	//funct.Test7()

	fmt.Println()
	err := funct.Monthwork("a")
	if err != nil {
		panic(any(err))
	}

	fmt.Println("主程序结束。。。。")
}
