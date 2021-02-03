package main

import (
	"errors"
	"fmt"
	"os"
)

// CircleQueue 环形队列
type CircleQueue struct {
	maxsize int    // 最大值4
	array   [5]int //数组
	head    int    // 队首 0
	tail    int    // 队尾 0
}

// Push 入队列
func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("queue full")
	}
	c.array[c.tail] = val

	// c.tail 在队列尾部,不包含最后元素
	// 队尾留一位,用作判断head是否为空,那队列的实际大小为maxsize -1
	c.tail = (c.tail + 1) % c.maxsize
	return nil
}

// Pop 出队列
func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	// 取出,head指向队首含队首元素
	c.array[c.head] = val
	c.head = (c.head + 1) % c.maxsize
	return
}

// ListQueue 显示队列
func (c *CircleQueue) ListQueue() {

	fmt.Println("环形队列情况如下:")
	// 取出当前队列有多少个元素
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	// 设计一个辅助变量.指向head
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxsize
	}
	fmt.Println()
}

// IsFull 判断是否为满
func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxsize == c.head
}

// IsEmpty 判断是否为空
func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

// Size 取出有多少个元素
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxsize - c.head) % c.maxsize
}

func main() {
	// 初始化队列
	queue := &CircleQueue{
		maxsize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int

	for {
		fmt.Println("1.add  添加队列")
		fmt.Println("2.get  获取队列")
		fmt.Println("3.show 查询队列")
		fmt.Println("4.exit 退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add", "1":
			fmt.Println("输入你要入对的列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get", "2":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列中数为", val)
			}
		case "show", "3":
			queue.ListQueue()
		case "exit", "4":
			os.Exit(0)
		}
	}
}
