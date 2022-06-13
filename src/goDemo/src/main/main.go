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
