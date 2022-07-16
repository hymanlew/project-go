package funct

import (
	"fmt"
	"goDemo/demo/pkg/utils"
)

var age = test()

//每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，被 Go 运行框架调用。因此通常可以在 init 函数中完成初始化工作。
//如果一个文件同时包含全局变量，init 函数和 main 函数，则执行流程是：全局变量定义 --> init 函数 --> main 函数。
func test() int {
	fmt.Println("global var...") //1
	return 20
}

func init() {
	fmt.Println("init()...") //2
}

func mainTest() {
	//如果源文件的引入包中，也有全局变量定义或 init 函数时，则就会先执行引入包中的变量及 init 函数
	fmt.Println("main()...age=", age) //3
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}

//匿名函数的使用
var (
	//fun1 就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

var Fun2 = func(n1 int, n2 int) int {
	return n1 + n2
}

func test4() {
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	//将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数，即可多次调用
	//此时 a 的数据类型就是函数类型，我们可以通过 a 完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(90, 30)
	fmt.Println("res3=", res3)

	//全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println("res4=", res4)
	res5 := Fun2(4, 9)
	fmt.Println("res4=", res5)
}
