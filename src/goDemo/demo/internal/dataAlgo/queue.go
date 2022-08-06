package dataAlgo

import (
	"errors"
	"fmt"
	"os"
)

//队列是一个有序列表，可以用数组或链表来实现。遵循先入先出的原则。

//1，使用单向数组实现此结构：
//其中 maxSize 为该队列最大容量，first 为队首，last 为队尾。两者会随着数据的变化而改变。
//插入数据时，将队尾往后移，first == last, last + 1。
//若队尾 last 小于等于 maxSize - 1 时，则代表可存入数据，放在 last 的位置。否则就无法存入。
//即数组第一个元素为队首，最后一个为队尾。可以实现插入读取的操作。
//但这有一个问题，当操作到最后一个元素时，队首队尾都这个元素，无法读取也无法插入。且之前的坐标也无法复用了。

// Queue 注意项：初始值 first = last = -1，first 当前坐标 +1 才是此坐标的数据，last 坐标指向的就是当前的数据。
type Queue struct {
	maxSize int
	array   [5]int // 数组=>模拟队列
	first   int    // 表示队列首坐标
	last    int    // 表示队列尾坐标
}

func (que *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if que.last == que.maxSize-1 {
		return errors.New("queue full")
	}

	//队尾坐标后移
	que.last++
	que.array[que.last] = val
	return
}

func (que *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if que.last == que.first {
		return -1, errors.New("queue empty")
	}
	que.first++
	val = que.array[que.first]
	return val, err
}

func (que *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")

	//que.first 不包含队首的元素
	for i := que.first + 1; i <= que.last; i++ {
		fmt.Printf("array[%d]=%d\t", i, que.array[i])
	}
	fmt.Println()
}

//使用单向数组实现队列功能
func queue1() {
	queue := &Queue{
		maxSize: 5,
		first:   -1,
		last:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 从队列获取数据")
		fmt.Println("3. 输入show 显示队列")
		fmt.Println("4. 输入exit 退出队列")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

//2，使用环形数组来实现此结构:
//将数组看作是一个环形的，并通过取模的方法来实现。初始化为 first = last = 0。数组大小 size-1 即为最大容量。
//队尾坐标的下一个坐标为队首时表示队列满了，即将队列容量空出一个作为约定。
//即 (last + 1) % max = first 为队满。而 last == first 为队空。

type CircleQueue struct {
	maxSize int    // 4
	array   [5]int // 数组
	head    int    //队首 0
	end     int    //队尾 0
}

func (queue *CircleQueue) Push(val int) (err error) {
	if queue.IsFull() {
		return errors.New("queue full")
	}

	//先把值给当前的尾部 end，之后尾部坐标加 1
	queue.array[queue.end] = val
	queue.end = (queue.end + 1) % queue.maxSize
	return
}

func (queue *CircleQueue) Pop() (val int, err error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	//先取出当前 head 的值，之后队首坐标加 1
	val = queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	return
}

func (queue *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")

	//取出当前队列有多少个元素
	size := queue.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	//设计一个辅助的变量，指向head
	tempHead := queue.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, queue.array[tempHead])
		tempHead = (tempHead + 1) % queue.maxSize
	}
	fmt.Println()
}

func (queue *CircleQueue) IsFull() bool {
	return (queue.end+1)%queue.maxSize == queue.head
}

func (queue *CircleQueue) IsEmpty() bool {
	return queue.end == queue.head
}

func (queue *CircleQueue) Size() int {
	return (queue.end + queue.maxSize - queue.head) % queue.maxSize
}

func queue2() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		end:     0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

//3，使用链表实现队列，链表本身就是有序的列表。
//单链表有一个头结点，里面还指向了后面一个结点的指针。头结点的作用主要是用于标识链表头，本身此结点不存放任何数据。
//而其他的每个结点都有一个指针指向后面一个节点。

type HeroNode struct {
	num      int
	name     string
	nickname string
	next     *HeroNode //指向下一个结点
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1. 先找到该链表的最后一个结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head
	for {
		//表示找到最后
		if temp.next == nil {
			break
		}
		//让temp不断的指向下一个结点
		temp = temp.next
	}

	//3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
}

// InsertHeroNode2 插入节点的第二种方法，根据 num 的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1. 找到适当的结点
	//2. 创建一个辅助结点
	temp := head
	flag := true

	//让插入结点的 num 和 temp 下一结点的 num 比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.num >= newHeroNode.num {
			//说明 newHeroNode 就应该插入到 temp 后面，即插入到 temp next 的前面
			break
		} else if temp.next.num == newHeroNode.num {
			//说明链表中已经有这个 num, 就不能插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，节点已经存在 num =", newHeroNode.num)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := false

	//找到要删除结点的 num，和 temp 下一个结点的 num 比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.num == id {
			flag = true
			break
		}
		temp = temp.next
	}

	//说明已经找到, 需要删除
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

func ListHeroNode(head *HeroNode) {
	//创建一个辅助结点
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return
	}

	//遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.num, temp.next.name, temp.next.nickname)
		temp = temp.next

		//判断是否链表后
		if temp.next == nil {
			break
		}
	}
}

func queue3() {
	//1. 先创建一个头结点,
	head := &HeroNode{}

	//2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		num:      1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		num:      2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		num:      3,
		name:     "林冲",
		nickname: "豹子头",
	}

	// hero4 := &HeroNode{
	// 	num : 3,
	// 	name : "吴用",
	// 	nickname : "智多星",
	// }

	//3. 加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)

	//4. 显示
	ListHeroNode(head)

	//5 删除
	fmt.Println()
	DelHerNode(head, 1)
	DelHerNode(head, 3)
	ListHeroNode(head)
}
