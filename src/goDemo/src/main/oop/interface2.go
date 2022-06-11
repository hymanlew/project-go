package oop

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// Len 实现系统排序的 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {

	//这里按 Hero 的年龄从小到大排序
	return hs[i].Age < hs[j].Age

	//修改成对 Name 排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp

	//下面的一句话等价于上面的三句话
	hs[i], hs[j] = hs[j], hs[i]
}

func test3() {
	//对 intSlice 切片进行排序，实现方法有
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对自定义结构体切片进行排序，实现方法有
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将 hero append到 heroes切片
		heroes = append(heroes, hero)
	}

	//排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("-----------排序后------------")

	//排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i=", i, "j=", j) // i=20 j = 10
}

// ============================== 分隔线 =======================================

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, " 生来会爬树..")
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey //继承
}

func (little *LittleMonkey) Flying() {
	fmt.Println(little.Name, " 通过学习，会飞翔...")
}

func (little *LittleMonkey) Swimming() {
	fmt.Println(little.Name, " 通过学习，会游泳..")
}

//- 当 A 结构体继承了 B 结构体，则 A 就自动继承了 B 结构体的字段和方法，并且可以直接使用。
//- 当 A 结构体需要扩展功能时，同时不希望破坏继承关系，则可以去实现某个接口即可。因此可以认为：实现接口是对继承机制的补充。
func test4() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
