package dataAlgo

import "fmt"

//4，使用双向链表（带 head 头的）实现队列。
//单向链表的缺点有：
//1，单向链表查找的方向只能是一个方向，而双向链表可以向前或向后查找。
//2，单向链表不能自我删除，需要靠辅助节点（只能找到下一个节点来帮助删除）。而双向链表可以自我删除。

type HeroNodeD struct {
	num      int
	name     string
	nickname string
	pre      *HeroNodeD //指向前一个结点
	next     *HeroNodeD //指向下一个结点
}

func InsertHeroNodeD(head *HeroNodeD, newHeroNode *HeroNodeD) {
	//1. 先找到该链表的最后这个结点
	//2. 创建一个辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func InsertHeroNodeD2(head *HeroNodeD, newHeroNode *HeroNodeD) {
	//1. 找到适当的结点
	//2. 创建一个辅助结点
	temp := head
	flag := true

	//让插入的结点的 no，和 temp 下一个结点的 no 比较
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.num >= newHeroNode.num {
			//说明 newHeroNode 就应该插入到 temp 后面
			break
		} else if temp.next.num == newHeroNode.num {
			//说明链表中已经有这个no, 就不能插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.num)
		return
	} else {
		newHeroNode.next = temp.next //ok
		newHeroNode.pre = temp       //ok
		if temp.next != nil {
			temp.next.pre = newHeroNode //ok
		}
		temp.next = newHeroNode //ok
	}
}

func DelHerNodeD(head *HeroNodeD, id int) {
	temp := head
	flag := false

	//找到要删除结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.num == id {
			//说明我们找到了.
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next //ok
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

func ListHeroNodeD(head *HeroNodeD) {
	//1. 创建一个辅助结点
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return
	}

	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.num, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListHeroNodeD2(head *HeroNodeD) {
	//1. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return
	}

	//2. 让 temp 定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.num, temp.name, temp.nickname)

		//判断是否链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func queue4() {
	//1. 先创建一个头结点,
	head := &HeroNodeD{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNodeD{
		num:      1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNodeD{
		num:      2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNodeD{
		num:      3,
		name:     "林冲",
		nickname: "豹子头",
	}

	InsertHeroNodeD(head, hero1)
	InsertHeroNodeD(head, hero2)
	InsertHeroNodeD(head, hero3)
	ListHeroNodeD(head)

	fmt.Println("逆序打印")
	ListHeroNodeD2(head)
}

//5，单向环形链表的使用：

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环形
		fmt.Println(newCatNode, "第一个节点加入到环形的链表")
		return
	}

	//定义一个临时变量，去定位到环形头结点的，前一结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")

	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	for {
		//不可以直接输出 temp，否则 go 程序会自动遍历其每个字段的值，即会遍历其 next 结点循环，从而形成死循环
		fmt.Printf("猫的信息为=[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// DelCatNode 需要将最新的头结点返回，因为需要用头结点进行遍历输出
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	preNode := head

	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	//如果只有一个结点，且是删除头结点时，则就代表是清空链表了
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}

	//将 preNode 定位到链表最后
	for {
		if preNode.next == head {
			break
		}
		preNode = preNode.next
	}

	flag := true
	for {
		//如果到最后一个节点，则停止循环。且此节点还没比较
		if temp.next == head {
			break
		}
		if temp.no == id {
			//说明删除的是头结点
			if temp == head {
				head = head.next
			}
			preNode.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}

		//移动到下一个节点
		temp = temp.next
		//移动到下一个节点，且 helper 一直处于 temp 的上一个节点，用于做删除时的双向关联
		preNode = preNode.next
	}

	//如果 flag 为真，则表示最后一个节点之前的节点都不匹配，没有删除
	if flag {
		if temp.no == id {
			preNode.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	}
	return head
}

func queue5() {

	//这里初始化一个环形链表的头结点
	//不带头结点的意思是，头结点也是要指向下一个数据的
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 30)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)

}

//6，单向环形链表的使用场景：
//Josephu 问题：设编号为 1，2，...n 个人围坐一圈，约定编号为 k（1<=k<=n），然后从 1 开始报数，数到 m 的人出列。其下一位再从 1 报数，
//数到 m 的人出列。依次类推，直到所有人都出列为止，由此产生一个出队编号的序列。
//由此，可使用一个不带头结点的循环链表（即头结点也是要指向下一个数据的）来处理该问题，先构成一个有 n 个结点的单循环链表，然后由 k 结点
//起从 1 开始计数，计到 m 时则对应结点从链表中删除。然后依次循环，直到最后一个结点从链表中删除。

type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针
}

// AddBoy 构成单向环形链表，num 表示小孩的个数，*Boy 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点

	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	//循环构建这个环形链表
	for i := 1; i <= num; i++ {
		newBoy := &Boy{
			No: i,
		}

		//第一个小孩比较特殊, 不需要动
		if i == 1 {
			first = newBoy
			curBoy = newBoy
			curBoy.Next = first
		} else {
			curBoy.Next = newBoy
			curBoy = newBoy
			curBoy.Next = first //构造环形链表
		}
	}
	return first
}

func ShowBoy(first *Boy) {

	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {

	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}

	//声明一个变量 lastBoy，让其执行到环形链表的最后一个小孩, 它在删除小孩时需要用到
	lastBoy := first
	for {
		if lastBoy.Next == first {
			break
		}
		lastBoy = lastBoy.Next
	}

	//将 first 移动到 startNo 处，后面删除小孩时，就以 first 为准
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		lastBoy = lastBoy.Next
	}

	//开始数 countNum, 用 first 节点跟进，然后就删除 first 指向的小孩
	fmt.Println("开始游戏....")
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			lastBoy = lastBoy.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)

		//删除first执行的小孩
		first = first.Next
		lastBoy.Next = first

		//判断如果 lastBoy == first, 则表示圈子中只有一个小孩了
		if lastBoy == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 \n", first.No)
}

func queue6() {

	first := AddBoy(500)
	ShowBoy(first)
	PlayGame(first, 20, 31)
}
