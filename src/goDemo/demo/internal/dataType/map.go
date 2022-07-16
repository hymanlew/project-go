package dataType

import (
	"fmt"
	"sort"
)

func map1() {
	//- golang 中 map 的 key 可以是多种类型，如 bool、数字、string、指针、channel、还可以是只包含前面几个类型的接口、结构体、数组。
	//但是 slice，map，function 不可以做 key，因为这几个没法用 == 来判断。
	//- value 的类型和 key 基本一样，通常为数字（(整数、浮点数）、string、map、struct。
	//- map 的声明是不会分配内存的，初始化使用 make 分配内存后才能赋值和使用。

	//第一种方式
	var a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	heroes["hero4"] = "林冲"
	fmt.Println("heroes=", heroes)

	//案例
	//比如存放3个学生信息, 每个学生有 name和sex 信息
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
}

func map2() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	//因为 no3 这个 key 已经存在，因此下面这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)

	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//演示map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有 no1 这个 key，值为%v\n", val)
	} else {
		fmt.Printf("没有 no1 key\n")
	}

	//如果要删除 map 的所有 key，但没有方法一次性删除，可以遍历一下 key，然后逐个删除。
	//或者直接 make 一个新的 map，让原来的成为垃圾，被 gc 回收。
	cities = make(map[string]string)
	fmt.Println(cities)

}

func mapFor() {
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	fmt.Println("cities 有", len(cities), " 对 key-value")

	//使用for-range遍历一个结构比较复杂的map
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}

func mapslice() {

	//切片的数据类型如果是 map，则称为 slice of map，即 map 切片，这样 map 的个数就可以动态变化了。

	//使用一个 map 来记录 monster 的信息 name 和 age, 也就是一个 monster对应一个 map, 并且妖怪的个数可以动态的增加=>map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	// 下面这个写法越界。
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "狐狸精"
	// 	monsters[2]["age"] = "300"
	// }

	//这里需要使用到切片的 append 函数，动态的增加 monster
	newMonster := map[string]string{
		"name": "新的妖怪~火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}

func mapsort() {

	//golang 中没有专门的方法对 map 的 key 进行排序，且它默认是无序的，也不是按照添加的顺序存放的，即每次遍历，得到的输出都可能不一样。
	//可以手动先将 key 进行排序，然后根据 key 值遍历输出即可。

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	//如果按照 map 的 key 的顺序进行排序输出
	//1. 先将 map 的 key 放入到切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照 key 来输出 map 的值

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)
	//忽略下标，只需要下标对应的值
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}

}

// Stu 定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func valueStruct() {
	//map的 value 也经常使用 struct 类型，更适合管理复杂的数据
	//1.map 的 key 为 学生的学号，是唯一的
	//2.map 的 value为结构体，包含学生的名字，年龄, 地址

	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)

	//遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.Age)
		fmt.Printf("学生的地址是%v \n", v.Address)
		fmt.Println()
	}
}

/*
1)使用 map[string]map[string]sting 的 map 类型
2)key: 表示用户名，是唯一的，不可以重复
3)如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息（包括昵称 nickname 和密码 pwd）。
4)编写一个函数 modifyUser(users map[string]map[string]sting, name string) 完成上述功能
*/
func modifyUser(users map[string]map[string]string, name string) {
	//判断 users 中是否有 name，返回的 ok 类型为 bool 类型
	//v , ok := users[name]
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name //示意
	}
}

func work4() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"

	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")

	fmt.Println(users)
}
