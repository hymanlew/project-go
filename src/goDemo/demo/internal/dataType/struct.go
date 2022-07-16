package dataType

import (
	"encoding/json"
	"fmt"
)

/*
- 结构体是自定义的数据类型，代表一类事物。
- 结构体变量（实例）是具体的，实际的，代表一个具体变量。

从概念或叫法上看：
- 结构体字段 = 属性 = field
- 字段是结构体的一个组成部分，一般是基本数据类型、数组、也可是引用类型。
- 在创建结构体变量后，如果没有给字段赋值，则都对应一个零值(默认值)，布尔是 false，数值是 0，字符串是 ""。
- 指针，slice，和 map 的零值都是 nil，即还没有分配空间。如果需要使用这种字段，需要先 make 才能使用。
- 不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个，结构体是值类型。
*/

// Cat 定义一个 Cat 结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name   string
	Age    int
	Scores [3]int            //数组 [0,0,0]
	ptr    *int              //指针 nil
	slice  []int             //切片 []
	catmap map[string]string //map map[]
}

type Monster struct {
	Name string
	Age  int
}

func demo() {
	// 张老太养了20只猫猫:
	// 一只名字叫小白,今年3岁,白色。
	// 还有一只叫小花,今年10岁,花色。

	// 创建一个Cat的变量
	var cat Cat
	fmt.Printf("cat1的地址=%p\n", &cat)
	fmt.Printf("cat1的值=%v\n", cat)

	// 以下条件全部为 true，即引用类型默认值都是 nil
	if cat.ptr == nil {
		fmt.Println("ok1")
	}
	if cat.slice == nil {
		fmt.Println("ok2")
	}
	if cat.catmap == nil {
		fmt.Println("ok3")
	}

	cat.Name = "小白"
	cat.Age = 3

	// new 主要用来分配值类型的内存，比如 int、float32,struct...，返回的是指针
	cat.ptr = new(int)
	cat.ptr = &cat.Age

	cat.slice = make([]int, 10)
	cat.slice[0] = 100 //ok

	// map 的容量达到后会自动扩容，并不会发生 panic
	cat.catmap = make(map[string]string)
	cat.catmap["key1"] = "tom~"
	fmt.Println("cat=", cat)

	//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个, 结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 //结构体是值类型，默认为值拷贝
	monster2.Name = "青牛精"

	fmt.Println("monster1=", monster1) //monster1= {牛魔王 500}
	fmt.Println("monster2=", monster2) //monster2= {青牛精 500}
}

func demo2() {
	var cat = Cat{
		"a",
		1,
		[3]int{0, 0, 0},
		nil,
		make([]int, 10),
		make(map[string]string),
	}
	cat.Name = "abc"
	fmt.Println(cat)

	var cat2 *Cat = new(Cat)
	(*cat2).Name = "abc"
	// 以下操作是由语言底层进行处理，会自动加上 *cat2 的取值操作
	cat2.Name = "abc"
	fmt.Println(*cat2)

	var cat3 = &Cat{
		"b",
		2,
		[3]int{1},
		nil,
		nil,
		nil,
	}
	(*cat3).Name = "abc2"
	// 以下操作是由语言底层进行处理，会自动加上 *cat2 的取值操作
	cat3.Name = "abc2"
	fmt.Println(*cat3)
}

type PointMap struct {
	x int
	y int
}

type RectMap struct {
	pointA, pointB PointMap
}

type RectMap2 struct {
	leftUp, rightDown *PointMap
}

func structVM() {
	//结构体中的所有字段在内存中是连续分布的，包括指针类型的字段，其本身地址也是连续的，但他们指向的地址不一定是连续的
	//（即实际值数据不一定连续）

	r1 := RectMap{PointMap{1, 2}, PointMap{3, 4}}
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.pointA.x, &r1.pointA.y, &r1.pointB.x, &r1.pointB.y)

	r2 := RectMap2{&PointMap{10, 20}, &PointMap{30, 40}}
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n", &r2.leftUp, &r2.rightDown)

	//他们指向的数据地址不一定是连续的...，这要看系统在运行时是如何分配
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n", r2.leftUp, r2.rightDown)
}

/*
struct 的每个字段上，可以写一个 tag，该 tag 可以通过反射机制获取，常见的场景就是序列化和反序列化。
- 之所以引用它，是因为其序列化之后，字段名首字母也是大写的，对于其它程序来说不方便。
- 而直接修改字段首字母为小写，则序列化之后就是空值，因为首字母小写就不能被外部访问。
- 因此需要在要序列化的字段上加上 tag 标签（使用波浪线下面的那个符号标注）。
*/
type Monster2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Montype Monster2

func instanceof() {
	//由于结构体是用户单独定义的类型，在和其它类型进行转换时，需要有完全相同的字段（名字、个数、类型）
	//即当 a = A(b) 可以转换时，就要求结构体的的字段要完全一样（包括名字、个数和类型）
	var a Monster
	var b Monster2
	a = Monster(b)
	fmt.Println(a, b)

	//当结构体进行 type 重新定义时（相当于取别名），则 Golang 认为是新的数据类型，不能直接相等赋值，但是相互间可以强转
	var c Montype
	//c = a
	c = Montype(a)
	fmt.Println(a, c)
}

func tojson() {
	var monster = Monster2{
		"牛魔王",
		500,
	}
	//该函数使用了反射机制，进行序列化，反序列化操作
	json, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json)
		fmt.Println(string(json))
	}
}
