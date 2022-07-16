package dataType

import "fmt"

func stringslice() {

	//string 底层是一个 byte 数组，因此 string 也可以进行切片处理
	str := "hello@atguigu"
	slice := str[6:]
	fmt.Println("slice=", slice)

	//string 是不可变的，就是说不能通过 str[0] = 'z' 方式来修改字符串
	//str[0] = 'z' [编译不会通过，报错，原因是string是不可变]

	//如果需要修改字符串，要先将 string --> []byte 或者 []rune --> 修改 --> 重写转成string
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//string 转成 []byte 后，可以处理英文和数字，但是不能处理中文
	//原因是 []byte 字节来处理 ，而一个汉字，是3个字节，因此就会出现乱码
	//解决方法是将 string 转成 []rune 即可，因为 []rune是按字符处理，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)
}

func fbn(n int) []uint64 {

	//声明一个切片，切片大小 n
	fbnSlice := make([]uint64, n)
	//第一个数和第二个数的斐波那契 为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//进行for循环来存放斐波那契的数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func work2() {
	/*
		1)可以接收一个 n int
		2)能够将斐波那契的数列放到切片中
		3)提示, 斐波那契的数列形式:
		arr[0] = 1; arr[1] = 1; arr[2]=2; arr[3] = 3; arr[4]=5; arr[5]=8

		思路
		1. 声明一个函数 fbn(n int) ([]uint64)
		2. 编程fbn(n int) 进行for循环来存放斐波那契的数列  0 =》 1 1 =》 1
	*/

	//测试一把看看是否好用
	fnbSlice := fbn(20)
	fmt.Println("fnbSlice=", fnbSlice)
}
