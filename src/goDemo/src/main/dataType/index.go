package dataType

import "fmt"

func Data()  {
	// 基本数据类型，其变量存的就是值，也叫值类型。获取变量的地址，用 &。
	// 指针类型，其变量存的是一个地址，这个地址指向的空间存的才是值。获取指针类型所指向的值，用 *。

	// 值类型（即基本数据类型），都有对应的指针类型，形式为 *数据类型。比如 \*int，\*float32。
	// 值类型包括：基本数据类型 int 系列，float 系列，bool，string，数组和结构体 struct。

	//基本数据类型在内存布局，输出内存地址
	var i int = 10
	fmt.Println("i的地址=", &i)


	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身所用的地址 &ptr
	//4. ptr 所指向的数据的地址 &i
	//4. ptr 所指向的真正数据 *ptr
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v", &ptr)
	fmt.Printf("ptr 指向的值=%v", *ptr)


	var num = 10
	fmt.Printf("num address = %v\n", &num)
	var point = &num
	fmt.Printf("point data = %v\n", *point)

	*point = 15
	fmt.Printf("num data = %v\n", num)
	fmt.Printf("%v = %v\n", point, &num)

	var str = "abc"
	var spoint = &str
	fmt.Printf("data = %v\n", *spoint)
}


//基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值
func test1(n1 int) {
	//函数中的变量是局部的，函数外不生效
	//var var1 = 1;
	n1 = n1 + 10
	fmt.Println("test1() n1= ", n1)
}

//如果希望函数内的变量能修改函数外的变量，即可以传入值数据类型变量的地址 &*，函数内以指针的方式操作变量。从效果上看类似引用
// n1 就是 *int 类型
func test2(n1 *int) {
	fmt.Printf("n1的地址 %v\n",&n1)
	*n1 = *n1 + 10
	fmt.Println("test03() n1= ", *n1) // 30
}

func test3() {
	// num := 20
	// test1(num)
	// fmt.Println("main() num= ", num)

	num := 20
	fmt.Printf("num的地址=%v\n", &num)
	test2(&num)
	fmt.Println("main() num= ", num) // 30
}