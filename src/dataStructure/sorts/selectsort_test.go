package sorts

import (
	"fmt"
	"testing"
)

func SelectSort(arr *[5]int) {
	// 先完成第一个最大值和arr[0]交换
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}

		// 交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 %v\n", j+1, arr)
	}
}

func TestSelectSort(t *testing.T) {
	arr := [5]int{10, 34, 19, 100, 80}
	SelectSort(&arr)
	t.Log(arr)
}
