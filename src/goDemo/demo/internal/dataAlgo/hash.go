package dataAlgo

import (
	"fmt"
	"os"
)

//散列表，又叫哈希表，是根据关键码值（key、value）通过映射函数进行直接访问的数据结构。
//这个映射函数叫散列函数，存放记录的数组+链表叫散列表。

//Emp 定义哈希中链表中的元素
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (emp *Emp) ShowMe() {
	fmt.Printf("链表 %d 找到该雇员 %d\n", emp.Id%7, emp.Id)
}

//EmpLink 定义哈希中的链表结构，不带表头, 即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//Insert 添加员工的方法, 保证添加时，编号从小到大
func (empLink *EmpLink) Insert(emp *Emp) {

	//声明辅助指针，其中 pre 在 cur 前面
	cur := empLink.Head
	var pre *Emp = cur

	//如果当前 EmpLink 是一个空链表，就直接放在头部即可
	if cur == nil {
		empLink.Head = emp
		return
	}

	//如果不是空链表，就给 emp 找到对应的位置并插入
	//思路是：让 cur 和 emp 比较，然后让 pre 保持在 cur 前面
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	//退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

func (empLink *EmpLink) ShowLink(valveNum int) {
	if empLink.Head == nil {
		fmt.Printf("链表 %d 为空\n", valveNum)
		return
	}

	//遍历当前的链表，并显示数据
	cur := empLink.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", valveNum, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (empLink *EmpLink) FindById(id int) *Emp {
	cur := empLink.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//HashTable 定义 hashtable 哈希表, 并且每个 key 都含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (hashTable *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := hashTable.HashFun(emp.Id)
	//使用对应的链表添加
	hashTable.LinkArr[linkNo].Insert(emp)
}

func (hashTable *HashTable) ShowAll() {
	for i := 0; i < len(hashTable.LinkArr); i++ {
		hashTable.LinkArr[i].ShowLink(i)
	}
}

func (hashTable *HashTable) HashFun(id int) int {
	//得到一个 hash 值，就是对应的链表的下标
	return id % 7
}

func (hashTable *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定将该雇员应该在哪个链表
	linkNo := hashTable.HashFun(id)
	return hashTable.LinkArr[linkNo].FindById(id)
}

func hashDemo() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable

	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)

		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)

		case "show":
			hashtable.ShowAll()

		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
