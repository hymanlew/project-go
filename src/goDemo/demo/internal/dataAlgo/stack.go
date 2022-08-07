package dataAlgo

import (
	"errors"
	"fmt"
	"strconv"
)

//栈（stack）是一个 FILO 的有序列表，即限制线性表中元素的插入和删除，只能在线性表的同一端里进行。
//而允许插入-删除的一端为栈顶 top，另一端为固定的一端为栈底 bottom。从而形成先入后出的有序列表。

// Stack 用数组模拟栈的实现及使用
type Stack struct {
	MaxTop int    //表示栈最大可以存放的个数
	Top    int    //表示栈顶
	arr    [5]int //使用数组模拟栈
}

func (stack *Stack) Push(val int) (err error) {

	//先判断栈是否满了，因为数组下标是从 0 开始，所以要减 1
	if stack.Top == stack.MaxTop-1 {
		return errors.New("stack full")
	}

	//放入数据，起初栈顶坐标为 -1
	stack.Top++
	stack.arr[stack.Top] = val
	return
}

func (stack *Stack) Pop() (val int, err error) {

	//判断栈是否空，，因为数组下标是从 0 开始，所以无数据时下标是 -1
	if stack.Top == -1 {
		return 0, errors.New("stack empty")
	}

	//先取值，再坐标--
	val = stack.arr[stack.Top]
	stack.Top--
	return val, nil
}

// List 遍历栈，注意需要从栈顶开始遍历
func (stack *Stack) List() {

	//先判断栈是否为空
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("栈的情况如下：")
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, stack.arr[i])
	}
}

func stack() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当栈顶为 -1 时，表示栈为空
	}

	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	//显示
	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val) // 5

	//显示
	stack.List()

	fmt.Println()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop() // 出错
	fmt.Println("出栈val=", val)

	//显示
	stack.List()
}

//栈的应用场景：
//1，子程序调用：在跳往子程序前，会先将本程序下个指令的地址存入栈中，直到子程序执行完后再取出，以回到原程序中.
//2，处理递归调用：和子程序调用类似，只是除了存入下一指令的地址外，还会将参数，变量等存入栈中。即每次递归都是独立的栈空间。
//3，表达式的转换与求值。
//4，二叉树的遍历。
//5，图形的深度优先（depth-first）搜索法。

//表达式计算，算法实现（3+2*6-5）：
//1，创建两个栈，numStack（存放数）、operStack（存放操作符）
//2，开始正向扫描算法表达式字符串，若第一个字符是数字，则直接入栈 numStack
//3，若字符是运算符，此时若 operStack 是空栈，则直接入栈，若此时有数据，则比较栈顶运算符与当前运算符。
//4，若栈顶值大于等于当前值，就从 operStack 栈中取出，并从 numStack 栈中也取出两个数字进行运算，最后将结果重新入栈。然后将下一个运算符直接入栈。
//5，若栈顶值小于当前值，则同样先计算入栈的优先运算符，最后将结果再入栈，而后将下一个运算符直接入栈。
//6，当扫描表达式完毕，依次从符号栈取出运算符，并从数栈中取出两个数，将计算后的结果再入数栈，直到符号栈为空。

// IsOper 判断是不是一个运算符
func (stack *Stack) IsOper(val int) bool {

	//字符对应的 ASCiI 编码（+, - , * , /）
	if val == 43 || val == 45 || val == 42 || val == 47 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误.")
	}
	return res
}

// Priority 返回某个运算符的优先级[* / => 1，+ - => 0]
func (stack *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func calcStack() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "30+30*6-4-6"

	//定义一个index ，帮助扫描exp
	index := 0
	//为了配合运算，定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		//依次取出字符，因为字符串底层就是一个字节切片，所以可以用下标取值
		ch := exp[index : index+1]

		//获取字符字节对应的 ASCiI 码，是一个数字
		//num := int(byte('i'))
		temp := int([]byte(ch)[0])

		//先检查是不是运算符
		if operStack.IsOper(temp) {

			//如果 operStack 是一个空栈，则直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//如果 opertStack 栈顶的运算符的优先级，大于等于当前准备入栈的运算符的优先级，则就从符号栈pop出，
				//并从数栈也 pop 两个数进行运算，运算后的结果再重新入数栈，当前符号再入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//将计算结果重新入数栈
					numStack.Push(result)
					//当前的符号再压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {

			//处理多位数的字符拼接，如 131 等多位数字
			keepNum += ch

			//检查是否已经到达算式的最后，如果已到最后，则直接将 keepNum 入数栈即可
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//如果不是到达算式的最后，则每次都要检查 index 后面一个字符是不是运算符，如果不是则代表还是数字，是需要拼接数字的，不能进行运算
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))

					//清空拼接的数字字符串
					keepNum = ""
				}
			}
		}

		//继续扫描，先判断 index 是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}

	//扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数运算结果，将结果入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//将计算结果重新入数栈
		numStack.Push(result)
	}

	//如果算法没有问题，表达式也是正确的，则结果就是 numStack 最后数
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}
