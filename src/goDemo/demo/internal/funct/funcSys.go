package funct

import "fmt"

func sys() {
	//- len：用来求长度，比如 string、array、slice、map、channel
	//- make：主要用来分配引用类型的内存，比如 channel、map、slice
	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v , num1的地址%v\n", num1, num1, &num1)

	//- new：主要用来分配值类型的内存，比如 int、float32,struct...，返回的是指针
	num2 := new(int) // *int
	//num2的类型%T => *int
	//num2的值 => 0xc04204c098 （这个地址是系统分配，是真实数据的空间地址）
	//num2的地址%v => 0xc04206a020  (这个地址是系统分配，是该指针自己的地址，应该是栈地址)

	*num2 = 100
	//num2指向的值 = 100
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址%v\n num2这个指针，指向的值=%v",
		num2, num2, &num2, *num2)
}
