package oop

import "fmt"

// Pupil 小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

// Pupil 结构体特有的方法
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

// Graduate 大学生
type Graduate struct {
	Student //嵌入了Student匿名结构体
}

// Graduate 结构体特有的方法
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.SetAge(10)
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.GetSum(1, 2))

	//子结构体对匿名结构体字段的访问可以简化，即可以直接方法匿名结构体的字段和方法
	//且当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
	graduate := &Graduate{}
	graduate.Name = "mary~"
	graduate.SetAge(12)
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
	fmt.Println("res=", graduate.Student.GetSum(10, 20))

	//结构体可以使用嵌套匿名结构体所有的字段和方法，首字母大写或者小写的字段、方法，都可以使用
	graduate.Student.testPrivateMethod()
}
