package process

import "fmt"

//在函数中，经常需要创建资源(比如：数据库连接、文件句柄、锁等)，为了在函数执行完毕后，及时的释放资源，Go 提供了 defer 延时机制。
//当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈（defer栈）, 然后继续执行函数下一个语句。
//在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。
//当函数执行完毕后，再从 defer 栈按照先入后出的方式出栈，执行程序

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1) //defer 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) //defer 2. ok2 n2= 20

	n1++                         // n1 = 11
	n2++                         // n2 = 21
	res := n1 + n2               // res = 32
	fmt.Println("ok3 res=", res) //defer 1. ok3 res= 32
	return res
}

func Test6() {
	res := sum(10, 20)
	fmt.Println("res=", res) // 4. res= 32
}

//defer 最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。
//在 golang 中的通常做法是，创建资源后(打开了文件，获取了数据库的链接，或者是锁资源)，可以执行 defer file.Close()/connect.Close()
//在 defer 后，可以继续使用创建资源。当函数完毕后，系统会依次从 defer 栈中，取出语句关闭资源。
//这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。
