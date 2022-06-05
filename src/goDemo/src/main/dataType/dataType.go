package dataType

import (
	"fmt"
	"unsafe"
)

func intT()  {
	var n1 = 100
	// %T 查看某个变量的数据类型
	// fmt.Printf() 用于做格式化输出
	fmt.Printf("n1 的 类型 %T \n", n1)

	// 查看某个变量的占用字节大小和数据类型 （使用较多）
	var n2 int64 = 10
	//unsafe.Sizeof(n1) 是 unsafe 包的一个函数，可以返回变量占用的字节数
	fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
}

func floatT()  {
	// 浮点型默认声明为 float64 类型
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T \n", num5)

	//十进制数形式：如：5.12   .512(必须要有小数点）
	num6 := 5.12
	num7 := .123 //=> 0.123
	fmt.Println("num6=", num6, "num7=", num7)

	//科学计数法形式
	num8 := 5.1234e2 // ? 5.1234 * 10的2次方
	num9 := 5.1234E2 // ? 5.1234 * 10的2次方 shift+alt+向下的箭头
	num10 := 5.1234E-2 // ? 5.1234 / 10的2次方 0.051234
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)

	// rune == int32
	var c1 rune = '北'
	fmt.Println("c1=", c1, unsafe.Sizeof(c1))
}

func char()  {
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//当直接输出 byte 值，就是输出了的对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '北' //overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

	//可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
	var c4 int = 22269 // 22269 -> '国' 120->'x'
	fmt.Printf("c4=%c \n", c4)

	//字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
	var n1 = 10 + 'a' //  10 + 97 = 107
	fmt.Println("n1=", n1)
}

func stringT() {
	// go 中的字符串属于基本数据类型。并且它是由字节 byte 组成，而不是类似 java 的由字符组成
	var address string = "北京长城 110 hello world!"
	fmt.Println(address)

	//字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的
	//var str = "hello"
	//str[0] = 'a' //这里就不能去修改str的内容，即go中的字符串是不可变的。

	//字符串的两种表示形式
	//	1, 双引号, 会识别转义字符
	str2 := "abc\nabc"
	fmt.Println(str2)

	//	2, 反引号, 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、 输出源代码等效果
	// 波浪号下面的符号
	str3 := ` 
	package main
	import (
		"fmt"
		"unsafe"
	)
	
	//演示golang中bool类型使用
	func main() {
		var b = false
		fmt.Println("b=", b)
		//注意事项
		//1. bool类型占用存储空间是1个字节
		fmt.Println("b 的占用空间 =", unsafe.Sizeof(b) )
		//2. bool类型只能取true或者false
		
	}
	`
	fmt.Println(str3)

	//字符串拼接方式
	var str = "hello " + "world"
	str += " haha!"

	fmt.Println(str)
	//当一个拼接的操作很长时，怎么办，可以分行写,但是注意，需要将+保留在上一行.
	str4 := "hello " + "world" + "hello " + "world" + "hello " +
		"world" + "hello " + "world" + "hello " + "world" +
		"hello " + "world"
	fmt.Println(str4)


	var a int // 0
	var b float32 // 0
	var c float64 // 0
	var isMarried bool // false
	var name string // ""
	//这里的 %v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v",a,b,c,isMarried, name)
}

func defaultvalue()  {
	// 在 go 中，数据类型都有一个默认值，当程序员没有赋值时，就会保留默认值，在 go 中，默认又叫零值。
	// 整型 = 0
	// 浮点型 = 0
	// 字符串 = ""
	// bool = false
}
