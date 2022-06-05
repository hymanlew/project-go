package process

import (
	"fmt"
)

// 自定义一个函数
func test(char byte) byte {
	return char + 1
}

func switchT() {
	//请编写一个程序，该程序可以接收一个字符，比如: a,b,c,d,e,f,g
	//a表示星期一，b表示星期二 …  根据用户的输入显示相依的信息.
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	//- 先执行表达式，得到值，然后和 case 的表达式进行比较，如果相等就匹配到，然后执行对应的 case 的语句块，然后退出 switch 控制。
	//- 如果 switch 表达式的值没有和任何的 case 的表达式匹配成功，则执行 default 的语句块。执行后退出 switch 的控制。
	//- case 后的表达式可以有多个（常量值、变量、有返回值的函数等都可以），使用逗号间隔。
	switch test(key) + 1 {
		case 'a':
			fmt.Println("周一, 猴子穿新衣")
		case 'b':
			fmt.Println("周二，猴子当小二")
		case 'c':
			fmt.Println("周三，猴子爬雪山")
		//...
		default:
			fmt.Println("输入有误...")
	}

	//case 后面可以有多个表达式
	//case 语句块不需要写 break，因为默认会有。即默认情况下，当程序执行完 case 语句块后，就直接退出该 switch 控制结构。
	var n1 int32 = 51
	var n2 int32 = 20
	switch n1 {
		case n2, 10, 5:
			fmt.Println("ok1")
		case 90:
			fmt.Println("ok2~")
	}

	//switch 后也可以不带表达式，类似 if --else分支来使用
	var age int = 10
	switch {
		case age == 10:
			fmt.Println("age == 10")
		case age == 20:
			fmt.Println("age == 20")
		default:
			fmt.Println("没有匹配到")
	}

	//case 中也可以对 范围进行判断
	var score int = 90
	switch {
		case score > 90:
			fmt.Println("成绩优秀..")
		case score >= 70 && score <= 90:
			fmt.Println("成绩优良...")
		case score >= 60 && score < 70:
			fmt.Println("成绩及格...")
		default:
			fmt.Println("不及格")
	}

	//switch 后也可以直接声明/定义一个变量，分号结束，不推荐
	switch grade := 90; {
		case grade > 90:
			fmt.Println("成绩优秀~..")
		case grade >= 70 && grade <= 90:
			fmt.Println("成绩优良~...")
		case grade >= 60 && grade < 70:
			fmt.Println("成绩及格~...")
		default:
			fmt.Println("不及格~")
	}

	//如果在 case 语句块后增加 fallthrough 语句，则会继续执行下一个 case，也叫 switch 穿透。
	var num int = 10
	switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough //默认只能穿透一层
		case 20:
			fmt.Println("ok2")
			fallthrough
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("没有匹配到..")
	}

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的变量类型。
	var x interface{}
	var y = 10.0
	x = y
	switch i:=x.(type) {
		case nil:
			fmt.Printf("x 的类型是 %T", i)
		case int:
			fmt.Printf("x 的类型是 int")
		case float32:
			fmt.Printf("x 的类型是 float32")
		case func(int) float64:
			fmt.Printf("x 的类型是 func(int)")
		case bool, string:
			fmt.Printf("x 的类型是 bool, string")
		default:
			fmt.Printf("x 的类型未知")
	}
}

func forT()  {
	//for循环快速入门
	for i := 1; i <= 10; i++ {
		fmt.Println("你好，尚硅谷", i)
	}

	//for循环的第二种写法
	j := 1 //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("你好，尚硅谷~", j)
		j++ //循环变量迭代
	}

	//for循环的第三种写法, 这种写法通常会配合 break 使用
	//此方式也等价 for ;; {}
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok~~", k)
		} else {
			break //break就是跳出这个for循环
		}
		k++
	}


	//如果字符串含有中文，则传统的遍历字符串方式就是错误的，会出现乱码。
	//原因是传统的对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节。此时就需要将 str 转成 []rune 切片。

	//字符串遍历方式1 - 传统方式
	// var str string = "hello,world!北京"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i]) //使用到下标...
	// }

	//字符串遍历方式1 - 传统方式
	var str string = "hello,world!北京"
	str2 := []rune(str) // 就是把 str 转成 []rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i]) //使用到下标...
	}

	fmt.Println()

	//字符串遍历方式2 - for-range
	//对应 for-range 遍历方式而言，是按照字符方式遍历的。因此如果字符串中有中文，也是 ok。
	str = "abc~ok上海"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}