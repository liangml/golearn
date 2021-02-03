package main

import (
	"errors"
	"fmt"
	"os"
)

// Queue 队列是一个有序列表，可以用"数组"或是"链表"来实现。
// 遵循先进先出的原则。
type Queue struct {
	maxsize int
	array   [5]int // 数组 => 模拟数组
	front   int    // 表示指向队列列首
	rear    int    // 表示指向队列的尾部
}

// AddQueue 添加数据到队列
func (c *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if c.rear == c.maxsize-1 { //rear是队列尾部(含最后一个元素)
		return errors.New("queue full")
	}
	fmt.Println("添加前", c.rear)
	c.rear++
	fmt.Println("添加后", c.rear, c.array)
	c.array[c.rear] = val
	return nil
}

// GetQueue 取出数据
func (c *Queue) GetQueue() (val int, err error) {
	// 判断队列是否为空
	if c.rear == c.front { // 队列为空
		return -1, errors.New("queue empty")
	}
	c.front++
	val = c.array[c.front]
	return val, err
}

// ShowQueue 显示队列,找到队首，遍历到最尾
func (c *Queue) ShowQueue() {

	fmt.Println("队列当前的情况:")
	// c.front 不包含队首到元素
	for i := c.front + 1; i <= c.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, c.array[i])
	}
}

func main() {
	queue := Queue{
		maxsize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get", "2":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列中数为", val)
			}
		case "show", "3":
			queue.ShowQueue()
		case "exit", "4":
			os.Exit(0)
		}
	}
}
