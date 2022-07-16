package dataType

func demo3() {
	//- 常量使用 const 修饰，且定义时必须初始化。
	//- 常量不可修改，只能修饰 bool，数值类型，string 类型。
	//- 常量名首字母小写是 private，首字母大写是 public。

	const name = "name"
	const num = 10
	const double = 9 / 3

	const doub2 = num / 3
	//num2 := 9
	//const doub3 = num2/3

	const (
		a = 10
		b = 20
	)

	const (
		//iota 就是常量值 0
		a1 = iota
		//b2 = a2 + 1
		b2
		//c2 = b2 + 1
		c2
		//当前面出现 iota 时，后面的 iota 就失效了，即还是会进行值递增
		d = iota
		e = iota
	)
}
