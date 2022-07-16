package method

import "fmt"

type person struct {
	Name string
	Age  int
	sex  int
}

//给 person 结构体添加 speak 方法, 输出xxx是一个好人
func (person person) speak() {
	person.Name = "jack"
	fmt.Println(person.Name, "是一个goodman~")
}

func (person person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func (person *person) speak2() string {
	person.Name = "jack"
	fmt.Println(person.Name, "是一个goodman~")
	return person.Name
}

func (person *person) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", person.Name, 10)
	return str
}

type Dog struct{}

/*
- Golang 中的方法是作用在指定的数据类型上的（即是和指定的数据类型绑定)，因此使用 type 自定义的类型，都可以有方法，而不仅仅是 struct。
- 即一个类型的方法是和该类型绑定的，且其方法只能通过该类型的实例变量调用，而不能直接调用，也不能使用其它类型变量来调用。

- 方法的访问控制的范围和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其它包访问。
- 如果一个类型实现了 String() 这个方法，则 fmt.Println 默认会调用本变量的 String() 方法进行输出。
*/
func demo() {
	var p person
	p.Name = "tom"

	//调用方法
	p.speak()
	//输出 tom，因为结构体是值传递的，在方法内部修改，不会影响到方法外部的数据
	fmt.Println(" p.Name=", p.Name)

	//下面的使用方式都是错误的
	// var dog Dog
	// dog.test()
	// test()

	//调用方法
	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("res=", res)

	fmt.Printf("person 结构体变量地址 =%p\n", &p)
	//name := (&p).speak2()
	//编译器底层做了优化 (&p).speak2() 等价 p.speak2(), 因为编译器会自动的给加上 &p
	name := p.speak2()
	fmt.Println("新名字 = ", name)

	//如果自定义实现了本类型的 String 方法，就会自动调用
	//并且不能设置和传入值类型参数，且 string 方法首字母要大写，否则不会调该方法
	fmt.Println(&p)
}

/*
- Golang 中的方法是作用在指定的数据类型上的（即是和指定的数据类型绑定)，因此使用 type 自定义的类型，都可以有方法，而不仅仅是 struct。
*/
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

func (i *integer) change() {
	*i = *i + 1
}

func demo2() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)
}

/*
工厂模式案例
- 如果结构体的首字母是大写的，则在其它包引入包并创建其实例，是没有问题的。可以直接创建结构体的变量(实例)。
- 但如果首字母是小写的，则就不能直接创建了，需要使用工厂模式来解决。
- 即使用工厂模式实现跨包创建结构体实例。
*/
func NewStudent(name string, age int, sex int) *person {
	return &person{
		Name: name,
		Age:  age,
		sex:  sex,
	}
}

//如果 sex 字段首字母小写，则在其它包不可以直接方法，可以提供一个方法
func (s *person) GetSex() int {
	return s.sex
}
