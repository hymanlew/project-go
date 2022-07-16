package oop

import "fmt"

// Usb 定义一个接口
type Usb interface {
	// Start Stop 声明两个没有实现的方法
	Start()
	Stop()
}

// Usb2 定义另一个接口
type Usb2 interface {
	Start()
	Stop()
}

//只要一个变量含有了接口类型中的所有方法，那么这个变量就是实现这个接口。即 Golang 中接口的实现，是基于方法去判断是否被实现，
//而不是基于接口名。即接口的实现是不需要去显示的声明，也没有 implement 这样的关键字

type Phone struct{}

// Start 让 Phone 实现 Usb接口，即就是实现了 Usb 接口声明所有方法。同时也是实现了 Usb2 接口
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct{}

// Start 让 Camera 实现 Usb 接口，同时也是实现了 Usb2 接口
func (c Camera) Start() {
	fmt.Println("相机开始工作~~~。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer struct{}

// Working 编写一个方法，接收一个 Usb 或 Usb2 接口类型变量
//- interface 类型默认是一个指针（引用类型），如果没有对 interface 初始化就使用，那么会输出 nil。
func (c Computer) Working(usb Usb) {

	//通过 usb 接口变量来调用 Start和Stop 方法
	usb.Start()
	usb.Stop()
}

func test1() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//- interface 类型默认是一个指针（引用类型），如果没有对 interface 初始化就使用，那么会输出 nil。
	computer.Working(phone)
	computer.Working(camera)
}

// ============================== 分隔线 =======================================

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type CInterface interface {
	AInterface
	BInterface
	demo()
}

type EmptyInterface interface {
}

//- 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给该接口类型。
//- 并且只要是自定义数据类型就可以实现接口，而不仅仅是结构体类型。
type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

// Monster 一个自定义类型可以实现多个接口，且 Golang 的接口中不能有任何变量。
type Monster struct{}

func (m *Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

// Stu 一个接口 A 可以继承多个别的接口 B、C，这时如果要实现 A 接口，也必须将 B、C 接口的方法也全部实现。
type Stu struct{}

func (stu Stu) Say() {
	fmt.Println("A SAY")
}

func (stu Stu) Hello() {
	fmt.Println("B SAY")
}

func (stu Stu) demo() {
	fmt.Println("C SAY")
}

func test2() {
	//- 接口本身不能创建实例，但可以指向一个实现了该接口的自定义类型的变量（实例）。
	//var a AInterface
	//a.Say()

	var i integer = 10
	var a AInterface = i
	a.Say() // integer Say i = 10

	//Monster实现了 AInterface 和 BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = &monster
	a2.Say()
	b2.Hello()

	var c CInterface = Stu{}
	c.demo()

	//- 空接口 interface{} 没有任何方法，所以所有的类型都自动实现了空接口，即可以把任何一个变量赋给空接口。
	var e EmptyInterface = Stu{}
	fmt.Println(e)

	var num = 10
	e = num
	fmt.Println(e)

	var e2 interface{} = Stu{}
	fmt.Println(e2)
}
