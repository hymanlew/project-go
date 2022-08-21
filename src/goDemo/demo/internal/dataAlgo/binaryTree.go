package dataAlgo

import "fmt"

//树有很多种，每个节点最多只能有两个子节点时，它被称为二叉树。下面有子节点的称为子节点，最下面一级节点由于没有
//子节点，所以称为叶节点。下面说明下前序、中序、后序遍历。

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//PreOrder 前序遍历：先输出 root 结点，再输出左子树，然后再输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//InfixOrder 中序遍历：先输出 root 的左子树，再输出 root 结点，最后输出 root 的右子树
func InfixOrder(node *Hero) {
	if node != nil {

		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//PostOrder 后序遍历：先输出 root 的左子树，再输出 root 的右子树，最后输出 root 结点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func treeDemo() {
	//构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1

	node10 := &Hero{
		No:   10,
		Name: "tom",
	}
	node12 := &Hero{
		No:   12,
		Name: "jack",
	}
	left1.Left = node10
	left1.Right = node12

	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2

	//前序
	//PreOrder(root)
	//中序
	//InfixOrder(root)
	//后序
	PostOrder(root)
}
