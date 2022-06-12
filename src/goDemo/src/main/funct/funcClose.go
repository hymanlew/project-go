package funct

import (
	"fmt"
	"strconv"
	"strings"
)

//累加器闭包，AddUpper 是一个函数，返回的数据类型是 fun (int) int，是一个匿名函数。
//但是这个匿名函数有引用到函数外的 n，因此这个匿名函数就和 n 形成一个整体，构成闭包。
func AddUpper() func(int) int {
	var n int = 10
	var s = ""
	return func(x int) int {
		s += "---" + strconv.Itoa(x)
		fmt.Println(s)

		n = n + x
		return n
	}
}

func Test5() {
	//可理解为：闭包就是一个类, 函数是方法，n 是成员变量。函数和此函数使用到的 n 构成闭包。
	//当反复调用 f 函数时，由于 n 是只初始化一次，因此每调用一次就进行累计。
	//要搞清楚闭包的关键，就是要分析出返回的函数它引用到哪些变量，因为函数和它引用到的变量共同构成闭包。
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
}

// 1)编写一个函数, 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg), 则返回文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//传统的实现方法，每次调用都要传入后缀名
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func Test6() {
	//返回一个闭包
	f2 := makeSuffix(".jpg")               //使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter"))   // winter.jgp
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg

	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter"))   // winter.jgp
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "bird.jpg")) // bird.jpg
}
