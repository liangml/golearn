package main

import (
	"fmt"
	"os"
)

// Emp 定义emp
type Emp struct {
	ID   int
	Name string
	Next *Emp
}

// 删除?

// ShowMe 显示自己
func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到雇员%d\n", e.ID%7, e.ID)
}

// EmpLink 定义EmpLink
// 不打表头的
type EmpLink struct {
	Head *Emp
}

// Insert 添加员工方法,保证添加时编号是从小到大
func (e *EmpLink) Insert(emp *Emp) {
	cur := e.Head      // 辅助指针
	var pre *Emp = nil // 辅助指针
	// 如果当前EmpLink就是一个空链表
	if cur == nil {
		e.Head = emp // 完成
		return
	}
	// 如果不是一个空链表,给emp找到对应的位置并插入
	// 让cur 和 emp比较，让pre保持在cur前面
	for {
		if cur != nil {
			// 比较
			if cur.ID > emp.ID {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时,是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

// ShowLink 显示链表信息
func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	// 遍历当前列表显示数据
	cur := e.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.ID, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// FindByID 根据ID查找对应的雇员,如果没有就返回nil
func (e *EmpLink) FindByID(id int) *Emp {
	cur := e.Head
	for {
		if cur != nil && cur.ID == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// HashTable 定义一个hashtable,含有一个链表数组
type HashTable struct {
	LinArr [7]EmpLink
}

// Insert 给hashtable添加insert雇员方法
func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数,将雇员添加到列表
	LinkNo := h.HashFun(emp.ID)
	// 使用对应的链表添加
	h.LinArr[LinkNo].Insert(emp)
}

// ShowAll 显示hashtalbe所有雇员
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinArr); i++ {
		h.LinArr[i].ShowLink(i)
	}
}

// HashFun 编写一个散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7 // 得到对应列表的下标
}

// FindByID 找到ID
func (h *HashTable) FindByID(id int) *Emp {
	// 确定雇员在哪个列表
	linkNo := h.HashFun(id)
	return h.LinArr[linkNo].FindByID(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	hashtable := HashTable{}
	for {
		fmt.Println("=========雇员菜单==========")
		fmt.Println("input 添加雇员")
		fmt.Println("show  查看雇员")
		fmt.Println("find  查找雇员")
		fmt.Println("exit  退出雇员系统")
		fmt.Println("请输入")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入id")
			fmt.Scanln(&id)
			fmt.Println("输入name")
			fmt.Scanln(&name)
			emp := &Emp{
				ID:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入ID号")
			fmt.Scanln(&id)
			emp := hashtable.FindByID(id)
			if emp == nil {
				fmt.Printf("id=%d的雇员不存在\n", id)
			} else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入为空")
		}
	}
}
