package method

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  int
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name) // jack
}

func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test03() =", p.Name) // mary
}

func demo3() {
	//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	p := Person{"tom", 10, 1}
	test01(p)
	test02(&p)

	//对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
	p.test03()
	fmt.Println("p.name=", p.Name) // tom

	//从形式上是传入地址，但是本质仍然是值拷贝
	(&p).test03()
	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test04()
	fmt.Println("main() p.name=", p.Name) // mary
	p.test04()                            // 等价 (&p).test04 , 从形式上是传入值类型，但是本质仍然是地址拷贝

	//即不管调用形式如何，真正决定是值拷贝还是地址拷贝，是看这个方法是和哪个类型绑定
	//- 如果是和值类型，比如 (p Person)，则是值拷贝
	//- 如果和指针类型，比如 (p \*Person)，则是地址拷贝
}
