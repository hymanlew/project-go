package method

import "fmt"

func factory()  {
	var stu = NewStudent("tom~", 9,1)
	if stu == nil {
		panic(any("创建失败..."))
	}

	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, " score=", stu.GetSex())
}
