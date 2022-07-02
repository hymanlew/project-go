package reflect

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

// GetSum 方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set 方法，接收四个值给 s 赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// Print 方法，显示 s 的值
func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

/*
- 不是所有 go 类型的 Type 值都能使用所有方法，请参见每个方法的文档获取使用限制。
- 在调用有分类限定的方法时，应先使用 Kind 方法获知类型的分类。调用该分类不支持的方法会导致运行时 panic。
- ValueOf 返回一个初始化为 i 接口保管的具体值的 Value。
- ValueOf(nil) 返回 Value 零值。Value 类型的零值表示不持有某个值。
- 零值的 IsValid 方法返回 false，其 Kind 方法返回 Invalid，而 String 方法返回 <invalid Value>，所有其它方法都会panic。绝大多数函数和方法都永远不返回 Value 零值。
- 如果某个函数/方法返回了非法的 Value，则它在文档中一定会有显式的说明具体情况。
- 如果某个 go 类型值可以安全的用于多线程并发操作，则它的 Value 表示也可以安全的用于并发。
*/

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	//获取 a 对应的类别
	//kd := typ.Kind()
	kd := val.Kind()

	//如果传入的不是 struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) //4

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))

		//获取到 struct 标签, 注意要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")

		//如果该字段有 tag 标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//method 函数返回该类型方法集中的第 i 个方法，默认按方法名排序对就到 i 值，i从 0 开始（ASCII码）。
	//i 不在 [0, NumMethod()) 范围内时，将导致 panic
	//Call 方法使用输入的参数 in 调用 v 持有的函数。入参出参都是 []reflect.value
	val.Method(1).Call(nil) //获取到第二个方法，即 print 方法，调用它

	//调用结构体的第 1 个方法 Method(0)
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //即调用 getsum 方法，传入的参数是 []reflect.Value, 返回 []reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value
}

func demo2() {
	var monster Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(monster)
}
