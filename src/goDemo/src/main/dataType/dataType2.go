package dataType

import "fmt"

func changeType()  {
	//Golang 和 java/c 不同，Go 在不同类型的变量之间赋值时需要显式转换。也就是说 Golang 中数据类型不能自动转换。
	//Go 中，数据类型的转换可以是从，表示范围小 -> 表示范围大，也可以 范围大 -> 范围小。
	//被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化！

	//表达式 T(v)  将值 v 转换为类型 T
	//T：就是数据类型，比如 int32，int64，float32 等等
	//v：就是需要转换的变量

	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i) //低精度->高精度
	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n", i ,n1, n2, n3)

	//被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", i) // int32

	//在转换中，比如将 int64  转成 int8 【-128---127】 ，编译时不会报错，
	//只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1) //
	fmt.Println("num2=", num2)


	// var n1 int32 = 12
	// var n2 int64
	// var n3 int8
	// n2 = int64(n1) + 20  //int32 ---> int64 错误
	// n3 = int8(n1) + 20  //int32 ---> int8 错误
	// fmt.Println("n2=", n2, "n3=", n3)

	// todo 即在数据计算时，编译器只会校验两边的类型是否一致，不会考虑是否溢出
	// var n1 int32 = 12
	// var n3 int8
	// var n4 int8
	// n4 = int8(n1) + 127  //【编译通过，但结果不是127+12 ,按溢出处理】
	// n3 = int8(n1) + 128 //【编译不过】
	// fmt.Println(n4)
}

func index()  {
	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num =" , num)
	var a_b int = 20
	fmt.Println(a_b)

	var int int = 30
	fmt.Println(int)
}

func changliang()  {
	//- 常量**使用const修改**
	//- 常量在**定义的时候，必须初始化**
	//- 常量**不能修改**
	//- 常量**只能修饰bool、数值类型（int，float系列）、string类型**
	//- 语法：**const identifier[type]= value**
}
