package calc

import "fmt"

//声明一个函数
func test() bool {
	fmt.Println("test....")
	return true
}

func compare() {
	var i int = 10

	if i < 9 && test() {
		fmt.Println("ok...")
	}
	if i > 9 || test() {
		fmt.Println("hello...")
	}
}
