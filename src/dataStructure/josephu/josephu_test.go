package josephu

import (
	"fmt"
	"testing"
)

// 小孩结构体
type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针
}

// 编写一个函数，构成单向环形链表
// num 表示小孩的个数
// *Boy 返回该环形链表第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	// 判断
	if num < 1 {
		fmt.Println("num then 0 >")
		return first
	}
	// 循环构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 分析构成循环链表需要辅助指针
		// 第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first // 形成循环
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first // 构成环形链表
		}
	}
	return first
}

// 显示单向环形链表
func ShowBoy(first *Boy) {
	// 处理环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}

	// 创建一个指针,帮助遍历:至少有一个小孩
	curBoy := first
	for {
		fmt.Printf("小孩编号:%d ->", curBoy.No)
		// 推出条件
		if curBoy.Next == first {
			break
		}
		// curBoy 移动到下一个
		curBoy = curBoy.Next
	}
	fmt.Println()
}

/*
设编号为1,2,...n的n个人围坐一圈,约定编号k(1<=k<=n)的人从1开始报数,数到m的那个人出列,	它的下一位从1开始报数,疏导m的人又出列,以此类推直到所有人出列为止,由此产生一个出队编号.
*/
// 分析思路
// 编写一个函数,PlayGame(First *Boy,StartNo int,CountNum int)
// 最后使用一个算法,按照要求,在环形列表中只有下一个人
func PlayGame(First *Boy, StartNo int, CountNum int) {
	// 空链表单独处理
	if First.Next == nil {
		fmt.Println("空链表,没有小孩")
		return
	}
	// 判断StartNo不能大于CountNum
	// 创建一个辅助指针，帮助删除指针
	tail := First
	// 让tail指向环形列表的最后一个小孩,在删除小孩时需要用到
	for {
		if tail.Next == First { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	// 让first移动到startNo[删除小孩以first为准]
	for i := 1; i <= StartNo-1; i++ {
		First = First.Next
		tail = tail.Next
	}
	// 开始数CountNum下，删除first指向的小孩
	for {
		// 开始数CountNum-1次
		for i := 1; i <= StartNo-1; i++ {
			First = First.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号%d出圈 ->\n", First.No)
		// 删除first指向的小孩
		First = First.Next
		tail.Next = First
		// 判断如果 tail == first,圈中只有一个小孩
		if tail == First {
			break
		}
	}
	fmt.Printf("最后编号小孩%d,出圈 ->\n", First.No)
}
func TestJoseHu(t *testing.T) {
	first := AddBoy(5)
	// 显示
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
