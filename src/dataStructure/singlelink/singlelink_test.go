package singlelink

import (
	"fmt"
	"testing"
)

// HeroNode 创建一个链表struct
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 表示指向下一个节点
}

// InsertHeroNode 给链表插入一个节点
// 第一种插入方式，在单链表的最后加入.
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到链表的最后节点
	// 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { // 表示找到最后
			break
		}
		temp = temp.next //让temp指向下一个节点
	}
	// 将newHeroNode加入链表的最后
	temp.next = newHeroNode
}

// ListHeroNode 显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	// 创建一个辅助节点
	temp := head
	// 先判断是否为d空链表
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	// 遍历链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		// 判断是否为队尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func TestSingleLink(t *testing.T) {
	// 先创建一个头节点
	head := &HeroNode{}

	// 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	// 加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	ListHeroNode(head)
}
