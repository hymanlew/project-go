package dataType

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

//上面的声明方式，也可以改成一次性声明
var (
	n3 = 300
	n4 = 900
	name2 = "mary"
)

func varOne()  {
	fmt.Println("hello world!")

	// golang的变量使用方式
	// 第一种：指定变量类型，若不赋值，则使用默认值
	// int 默认值为 0
	var n int
	fmt.Println("变量 n =", n)
	n = 10
	fmt.Println("变量 n =", n)

	// 第二种：根据值自行判定变量类型(类型推导)
	var num  = 10.11
	fmt.Println("num=", num)

	// 第三种：省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	// := 的 :不能省略，否则错误
	// 下面方式等价于 var name string   name = "tom"
	name := "tom"
	fmt.Println("name=", name)
}

func varMany()  {
	// 一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1, "n2=", n2, "n3=", n3)

	// var n1, name , n3 = 100, "tom", 888
	// fmt.Println("n1=",n1, "name=", name, "n3=", n3)

	// 一次性声明多个变量, 同样可以使用类型推导
	// n1, name , n3 := 100, "tom~", 888
	// fmt.Println("n1=",n1, "name=", name, "n3=", n3)

	//输出全局变量
	//fmt.Println("n1=",n1, "name=", name, "n2=", n2)
	fmt.Println("n3=", n3, "name2=", name2, "n4=", n4)
}
