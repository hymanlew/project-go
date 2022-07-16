package oop

import "fmt"

type Xiaohua struct {
	Student //嵌入了Student匿名结构体
	Pupil
}

type Xiaoqing struct {
	s     Student //嵌入了Student匿名结构体
	p     Pupil
	int   //嵌入匿名基本数据类型
	score int
}

type Xiaoai struct {
	*Student //嵌入了Student匿名结构体
	*Pupil
}

func test() {

	//当一个结构体嵌入两个（或多个）匿名结构体时，如两个匿名结构体有相同的字段和方法（同时当前结构体本身没有同名的字段和方法）时，
	//则在访问时，就必须明确指定匿名结构体名字，否则编译报错。
	//var xiaohua = Xiaohua{
	//	"xiaohua",
	//	12,
	//	0,
	//	80,
	//}

	var xiaohua = Xiaohua{
		Student{
			"xiaohua",
			12,
			0,
			80,
		},
		Pupil{},
	}
	fmt.Println(xiaohua)

	//如果一个 struct 嵌套了一个有名结构体，则这种模式就是组合，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字，否则编译报错
	var xiaoqing Xiaoqing
	xiaoqing.s.Name = "xiaoqing-s"
	xiaoqing.p.Name = "xiaoqing-p"
	//使用匿名基本数据类型
	xiaoqing.int = 10
	xiaoqing.score = 90

	var xiaoai = Xiaoai{
		&Student{
			"xiaohua",
			12,
			0,
			80,
		},
		&Pupil{},
	}
	xiaoai.Pupil.SetAge(15)
	fmt.Println(*xiaoai.Student, *xiaoai.Pupil)

}
