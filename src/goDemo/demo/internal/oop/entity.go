package oop

import "fmt"

type Student struct {
	Name  string
	age   int
	sex   int
	Score int
}

//封装（encapsulation）就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只有通过被授权的操作或方法，
//才能对字段进行操作。

//封装的实现步骤
//- 将结构体、字段或属性的首字母小写。即其它包不能使用，类似 private。
//- 给结构体所在包提供一个工厂模式的函数，首字母大写。类似一个构造函数。
//
//- 提供一个首字母大写的 Set 方法，类似 public，用于对属性判断并赋值
//  func (var 结构体类型名) SetXxx(参数列表) (返回值列表) {
//      -- 加入数据验证的业务逻辑
//      var.字段 = 参数
//  }
//
//- 提供一个首字母大写的 Get 方法，类似 public，用于获取属性的值
//  func (var 结构体类型名) GetXxx() {
//      return var.age;
//  }
//
//- 在 Golang 开发中并没有特别强调封装，这点并不像 Java。不用总是用 java 的语法特性来看待 Golang，
//Golang 本身对面向对象的特性是做了简化的。

// NewStudent 工厂模式的函数，类似于构造器
func NewStudent(name string, age int, sex int) *Student {

	if age > 100 || age < 0 {
		fmt.Println("年龄不合法")
		return nil
	}
	if sex != 1 && sex != 0 {
		fmt.Println("性别不合法")
		return nil
	}

	return &Student{
		Name: name,
		age:  age,
		sex:  sex,
	}
}

// GetSex 如果 sex 字段首字母小写，则在其它包不可以直接方法，可以提供一个方法
func (stu *Student) GetSex() int {
	return stu.sex
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.age, stu.Score)
}

func (stu *Student) SetAge(age int) {
	stu.age = age
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (stu *Student) testPrivateMethod() {
	fmt.Println("方法名小写的方法")
}
