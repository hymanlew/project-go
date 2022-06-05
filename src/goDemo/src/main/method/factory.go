package method

import "fmt"

/*
Golang 的结构体没有构造函数，通常是使用工厂模式来解决实例化问题。
- 如果结构体的首字母是大写的，则在其它包引入包并创建其实例，是没有问题的。可以直接创建结构体的变量(实例)。
- 但如果首字母是小写的，则就不能直接创建了，需要使用工厂模式来解决。
- 即使用工厂模式实现跨包创建结构体实例。
 */
func factory()  {
	var stu = NewStudent("tom~", 9,1)

	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, " score=", stu.GetSex())
}
