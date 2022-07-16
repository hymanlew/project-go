package oop

import "fmt"

type Point struct {
	x int
	y int
}

func test5() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //oK

	// 如何将 a 赋给一个 Point 变量?
	var b Point
	// b = a 不可以直接转换
	// b = a.(Point) // 可以，使用类型断言，成功就赋值，失败就报错
	b = a.(Point)
	fmt.Println(b)

	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2  //空接口，可以接收任意类型
	// // x=>float32 [使用类型断言]
	// y := x.(float32)
	// fmt.Printf("y 的类型是 %T 值是=%v", y, y)

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2 //空接口，可以接收任意类型
	// x=>float32 [使用类型断言]

	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}

// TypeJudge 编写一个函数，判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {

	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, x)
		}
	}
}

func test6() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name = "tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
