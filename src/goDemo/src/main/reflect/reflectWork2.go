package reflect

import (
	"fmt"
	"reflect"
)

func Demo3() {
	call := func(n1 int, n2 int) {
		fmt.Println(n1, n2)
	}
	call2 := func(n1 int, n2 int, s string) {
		fmt.Println(n1, n2, s)
	}

	var (
		function reflect.Value
		intvalue []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		intvalue = make([]reflect.Value, n)

		for i := 0; i < n; i++ {
			intvalue[i] = reflect.ValueOf(args[i])
		}

		function = reflect.ValueOf(call)
		function.Call(intvalue)
	}

	bridge(call, 1, 2)
	bridge(call2, 3, 4, "abc")
}

type User struct {
	//在使用反射设置字段或方法时，首字母必须大写，表示为公开的。否则反射获取不到
	Id   int64
	Name string
}

func Demo4() {
	var (
		model *User
		utype reflect.Type
		value reflect.Value
	)
	model = &User{}
	utype = reflect.TypeOf(model)
	value = reflect.ValueOf(model)

	//输出为 ptr
	fmt.Println(value.Kind().String())

	//输出为 struct
	userValue := value.Elem()
	fmt.Println(userValue.Kind().String())

	//在使用反射设置字段或方法时，首字母必须大写，表示为公开的。否则反射获取不到
	v := userValue.FieldByName("id")
	fmt.Println(v)

	userValue.FieldByName("Id").SetInt(10)
	userValue.FieldByName("Name").SetString("abc")
	fmt.Println(*model)

	fmt.Println("============== 分隔线 =============")

	//该方法返回一个 value 类型值，它持有一个指向类型为 type 的新申请的零值的指针
	initUser := reflect.New(utype.Elem())
	//输出为 ptr
	fmt.Println(initUser.Kind().String())
	//输出为 struct
	fmt.Println(initUser.Elem().Kind().String())

	//model = a.(*User)
	//a := initUser.Interface()
	model = initUser.Interface().(*User)
	userValue = initUser.Elem()
	userValue.FieldByName("Id").SetInt(20)
	userValue.FieldByName("Name").SetString("ddd")
	fmt.Println(*model)

	modelstruct := userValue.Interface()
	u, ok := modelstruct.(User)
	if ok {
		u.Id = 30
		u.Name = "eee"
		fmt.Printf("未修改之前的值 = %v\n", modelstruct)
		fmt.Println(u)
	}
}
