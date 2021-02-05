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

// 删除
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("空的环形列表,无法删除")
		return head
	}

	// 如果只有一个节点
	if temp.next == head { // 只有一个节点
		temp.next = nil
		return head
	}

	// 将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	// 两个以上的节点
	flag := true
	for {
		if temp.next == head { // 说明已经到最后一个(最后一个没有比较)
			break
		}
		if temp.no == id {
			if temp == head { // 说明删除的是头节点
				head = head.next
			}
			// 找到了,直接删除
			helper.next = temp.next
			fmt.Printf("猫%d被删除\n", id)
			flag = false
			break
		}
		temp = temp.next     // 移动  比较是不是
		helper = helper.next // 移动,一旦找到删除节点可以删除

	}
	// 还需要最后一次比较
	if flag { // 如果flag为真,以上没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫%d被删除\n", id)
		} else {
			fmt.Printf("sorry no id %d\n", id)
		}
	}
	return head
}

func TestCircleSingleLink(t *testing.T) {
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
