package funct

import "fmt"

/*
- Go 函数支持返回多个值，这一点是其它编程语言没有的。
- 在接收返回值时，希望忽略某个返回值，则使用 “_” 符号表示占位忽略。
- 如果返回值只有一个，则（返回值类型列表）可以不写 ()。
 */
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func test1() {
	//调用 getSumAndSub
	//res1 = 3 res2 = -1
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

	//希望忽略某个返回值，则使用 _ 符号表示占位忽略
	_, res3 := getSumAndSub(3, 9)
	fmt.Println("res3=", res3)
}


/*
- 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！
- 以后每天猴子都吃其中的一半，然后再多吃一个。
- 当到第十天时，想再吃时（还没吃），发现只有1个桃子了。
- 问题：最初共多少个桃子？

1)第10天只有一个桃子
2)第9天有几个桃子 = (第10天桃子数量 + 1) * 2
3)规律: 第n天的桃子数据 peach(n) = (peach(n+1) + 1) * 2
*/
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n + 1) + 1) * 2
	}
}

func test2() {
	//1534
	fmt.Println("第1天桃子数量是=", peach(1))
}


//在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int ) int {
	return funvar(num1, num2)
}

//自定义数据类型，此时 myFun 就是 func(int, int) int类型
type myFunType func(int, int) int

//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun2(funvar myFunType, num1 int, num2 int ) int {
	return funvar(num1, num2)
}

func getSumAndSub2(n1 int, n2 int) (sub int, sum int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//其中，args 是 slice 切片（动态数组），通过 args[index] 可以访问到各个值
//如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后
func sum(n1 int, args... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func test3() {
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)

	res := a(10, 40) // 等价 res := getSum(10, 40)
	fmt.Println("res=", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)


	//给 int 取别名，在go中 myInt 和 int 虽然都是int类型，但是 go 认为 myInt和int 是两个类型
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) //注意这里依然需要显示转换，因为 go 认为 myInt和int 是两个类型
	fmt.Println("num1=", num1, "num2=",num2)


	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)


	a1, b1 := getSumAndSub2(1, 2)
	fmt.Printf("a=%v b=%v\n", a1, b1)


	res4 := sum(10, 0, -1, 90, 10,100)
	fmt.Println("res4=", res4)

}