package circlesinglelink

import (
	"fmt"
	"testing"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是否为第一支猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Printf("加入:%+v\n", newCatNode)
		return
	}

	// 定义一个临时变量
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出环形列表
func ListCircleLink(head *CatNode) {
	temp := head
	if head.next == nil {
		fmt.Println("空列表")
	}
	for {
		fmt.Printf("猫的信息为=[no=%d, name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func TestCircleSingleLink(t *testing.T) {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	InsertCatNode(head, cat1)
	ListCircleLink(head)
}
