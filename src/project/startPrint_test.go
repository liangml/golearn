package project

import (
	"fmt"
	"strings"
	"testing"
)

// 空心金字塔打印
func TestStartPrint(t *testing.T) {

	var totalLevel int = 9

	for i := 1; i <= totalLevel; i++ {

		// 打印前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示打印多少个*
		// 当前层2*当前层-1
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 打印99乘法表
func TestSumPrint(t *testing.T) {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}

// 递归函数
func TestRecursion(t *testing.T) {
	recurlsion(4)
}
func recurlsion(n int) {
	if n > 2 {
		n--
		recurlsion(n)
	} else {
		fmt.Printf("n = %v\n", n)
	}
}

// 递归方式求出斐波那契数 1,1,2,3,5,8,13....
func TestFib(t *testing.T) {
	t.Log(Fib(6))
}
func Fib(n int) int {
	if (n == 1) || (n == 2) {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

// 闭包
// 闭包可以保留传入的变量
func makeSuffix(suffix string) func(string) string {
	// 匿名函数与 外部suffix 构成闭包
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
func TestMakeSuffix(t *testing.T) {
	s := makeSuffix(".jpg")
	t.Log(s("liangml"))
	t.Log(s("bird"))
}

// 冒泡排序
func BubboSort(arr *[]int) {
	fmt.Println("排序前", arr)
	// 循环总列表次数-1
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交换
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
func TestBubbleSort(t *testing.T) {
	arr := []int{24, 69, 80, 57, 13}
	BubboSort(&arr)
	fmt.Println("排序后", &arr)
}

// 顺序查找
func SequentialSearch(name string) {
	names := []string{"哪吒", "如来", "孙悟空", "唐僧"}
	// 第一种方式
	for k, v := range names {
		if names[k] == name {
			fmt.Printf("找到%v,下标为%v\n", k, v)
			break
		} else {
			fmt.Println("Unknow")
		}
	}
	// 第二种方式
	index := -1
	for k, _ := range names {
		if names[k] == name {
			index = k
		}
	}
	if index != -1 {
		fmt.Printf("找到%v\n", index)
	}
}

func TestSequentialSearch(*testing.T) {
	SequentialSearch("唐僧")
}

// 二分查找:(有序数组)
// 先找出中间下标：middle = (leftindex + rightindex)/2 ,让中间下标值与指定值进行比较
// 使用递归实现
func BinaryFind(arr *[]int, leftindex int, rightindex int, findVal int) {

	// 如果左边数组下标大于右边则没有
	if leftindex > rightindex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间下标
	middle := (leftindex + rightindex) / 2

	if (*arr)[middle] > findVal {
		// 说明查找的数在左边数组
		BinaryFind(arr, leftindex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 说明查找的数在右边数组
		BinaryFind(arr, middle+1, rightindex, findVal)
	} else {
		// 找到
		fmt.Printf("找到了,下标为%v \n", middle)
	}
}
func TestBinaryFind(t *testing.T) {
	arr := []int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 89)
}

// 二维数组
// 声明、定义、赋值
func TestMultiArray(t *testing.T) {
	/*
		000000
		001000
		020300
		000000
	*/
	arr := [4][6]int{}
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println(arr)

	for k, _ := range arr {
		// fmt.Println(arr[k])
		for j, _ := range arr[k] {
			fmt.Print(arr[k][j], " ")
		}
		fmt.Println()
	}

	// 地址分布
	fmt.Printf("arr[0]的地址:%p\n", &arr[0])
	fmt.Printf("arr[0][0]的地址:%p\n", &arr[0][0])
	fmt.Printf("arr[1]的地址:%p\n", &arr[1])
	fmt.Printf("arr[1][1]的地址:%p\n", &arr[1][1])

}
