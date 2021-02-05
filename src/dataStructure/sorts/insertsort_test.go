package sorts

import (
	"fmt"
	"testing"
)

func InsertSort(arr *[6]int) {
	for i := 1; i < len(arr); i++ {
		// 第一次,给第二个元素找到合适的位置并插入
		insertVal := arr[i]
		insertIndex := i - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d插入后 %v \n", i, arr)
	}
}

func TestInsertSort(t *testing.T) {
	arr := [...]int{23, 0, 12, 56, 34, 66}
	t.Log("排序前", arr)
	InsertSort(&arr)
	t.Log("排序后", arr)
}
