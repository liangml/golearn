package sorts

import (
	"fmt"
	"testing"
)

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
