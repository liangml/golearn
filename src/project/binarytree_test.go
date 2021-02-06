package project

import (
	"fmt"
	"hash"
	"testing"
)

type Hrs struct {
	No    int
	Name  string
	Left  *Hrs
	Right *Hrs
}

// 二叉树前序遍历
// 先输出root节点,再输出左子树,再输出右子树
func PreOrder(node *Hrs) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历
// 先输出root 左子树,再输出root节点,最后输出右子树
func InfixOrder(node *Hrs) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Right)
	}
}

// 后序遍历:将左边全部打完,把右边全部打完,打印root
func PostOrder(node *Hrs) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func TestBinaryTree(t *testing.T) {
	// 构建二叉树
	root := &Hrs{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hrs{
		No:   2,
		Name: "吴用",
	}
	node10 := &Hrs{
		No:   10,
		Name: "tom",
	}
	node12 := &Hrs{
		No:   12,
		Name: "jack",
	}
	left1.Left = node10
	left1.Right = node12
	right1 := &Hrs{
		No:   3,
		Name: "卢俊义",
	}
	right2 := &Hrs{
		No:   4,
		Name: "林冲",
	}

	root.Left = left1
	root.Right = right1
	right1.Right = right2
	// PreOrder(root)
	// InfixOrder(root)
	PostOrder(root)
}
