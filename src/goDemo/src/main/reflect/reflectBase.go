package reflect

import (
	"fmt"
	"reflect"
)

/*
- 反射可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别（kind)。
- 如果是结构体变量，还可以获取到结构体本身的字段、方法等信息。
- 通过反射，可以修改变量的值，可以调用关联的方法。

reflect 包实现了运行时反射，允许程序操作任意类型的对象。典型用法是：
- 用静态类型 interface{} 保存一个值，即其反射类是一个接口，通过该接口调用所有反射的方法。
- 调用 relect.TypeOf 函数获取其动态类型信息，返回一个 Type 类型值，返回接口中保存的值的类型，TypeOf(nil) 会返回 nil。是一个接口。
- 调用 ValueOf 函数返回一个 Value 类型的值，是一个结构体，该值代表运行时的数据。
- Zero 接受一个 Type 类型参数并返回一个代表该类型零值的 Value 类型值。

变量、interface{} 和 reflect.Value 是可以相互转换的，这在实际开发中，会经常使用到。
- 即任意的一个对象都是实现了一个空接口。
- 而任意一个接口都可以通过 ValueOf 函数转换为 reflect.Value 类型。
- 然后可通过 value.interface 函数，将 value 重新转换为一个空接口。
- 最后使用类型断言将空接口转换为最初的数据类型，value.(realType)。
*/

//通过反射获取传入变量的 type, kind 值
func reflectTest(b interface{}) {

	//先获取到 reflect.Type，代表值的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//获取到 reflect.Value，代表运行时的数据
	rVal := reflect.ValueOf(b)

	//n2 := 2 + rVal
	//type = reflect.Value, 不能直接参与数据计算处理的
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//将 rVal 转成 interface{}
	iV := rVal.Interface()

	//将 interface{} 通过断言转成需要的类型
	num2, ok := iV.(int)
	if ok {
		fmt.Println(num2 + 10)
		n2 := rVal.Int() + 10
		fmt.Println(n2)
	}

	//该方式只有用于 switch 语句中做断言
	//if iV.(type) == int {}
}

/*
- Type 是类型，Kind 是类别（即 go 内置的数据类型）。
- Kind 代表了 Type 类型值表示的具体分类。零值表示非法分类。

- Type 和 Kind 可能是相同的，也可能是不同的。
- 比如 int 类型的值，其 Type 是 int，Kind 也是 int。
- 比如结构体 Student，其 Type 是 pakege.Student，Kind 是 struct 即结构体类型。
*/
func reflectTest2(b interface{}) {

	//先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	//获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//获取变量对应的 Kind，即 go 内置的数据类型
	kind2 := rTyp.Kind()
	kind1 := rVal.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	//将 rVal 转成 interface{}, type 就是传入值的实际类型，iv 就是实际的值
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

func reflect01(b interface{}) {

	//kind = 指针，因为传入的 b 是一个地址
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())

	//通过反射来修改变量，当使用 SetXxx 方法设置值时，必须通过对应的指针类型来完成，这样才能改变传入变量的值, 否则报错
	//rVal.SetInt(20)

	//Elem 函数返回 v 持有的接口保管的值的 Value 封装，或 v 持有的指针指向的值的 Value 封装
	//因为只有 value 才能调用 set 方法，而对于指针类型则必须要获取的直接的值的 value 封装
	rVal.Elem().SetInt(20)

	//rVal.Elem() 可以这样理解
	// num := 9
	// ptr *int = &num
	// num2 := *ptr === rVal.Elem()
}

type Student struct {
	Name string
	Age  int
}

func demo() {
	//反射 type = int, value = 100
	var num int = 100
	reflectTest(num)

	//反射 type = reflectBase.Student, value = json内容字符串, kind = struct
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest2(stu)

	//通过反射, 修改 num int 的值, 修改 student 的值
	reflect01(&num)
	fmt.Println("num=", num) // 20
}
