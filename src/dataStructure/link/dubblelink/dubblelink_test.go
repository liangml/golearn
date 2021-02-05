package dubblelink

import (
	"fmt"
	"testing"
)

// HeroNode 创建一个链表struct
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode // 表示指向下一个节点
}

// InsertHeroNode 给双向链表插入一个节点
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
	newHeroNode.pre = temp
}

// InsertHeroNode2 第二种方法，根据no 的编号从小到大插入.
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 找到适当的节点
	// 创建一个辅助节点
	temp := head
	flag := true
	// 让插入的节点no 与temp的下一个节点no比较
	for {
		if temp.next == nil { // 说明到链表最后
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明newHeroNode就应该插入temp的后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明no已经在链表中，不允许插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("sorry , existing", newHeroNode.no)
		return
	} else {
		// 现将newHeroNode与temp的下一个结构绑定
		// fmt.Printf("Head: %+v\n", head)
		// fmt.Printf("Temp: %+v\n", temp)
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		// fmt.Printf("Banding newHeroNode = temp.next :%+v\n", newHeroNode.next)
		// 将temp的机构与新结构newHeroNode完成链绑定
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
		// fmt.Printf("Banding temp.next = newHeroNode :%+v\n", temp.next)
		// fmt.Printf("Head: %+v\n", head)
	}
}

// 删除一个节点
func DelHeroNode(head *HeroNode, id int) {
	// 创建一个辅助节点
	temp := head
	flag := false
	// 让删除的节点no 与temp的下一个节点no比较
	for {
		if temp.next == nil { // 说明到链表最后
			break
		} else if temp.next.no == id {
			// 说明已经找到
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		// 将temp.next指向下下个next,即跳过中间next
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除ID不存在")
	}
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
	fmt.Println()
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
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	ListHeroNode(head)
}
