package sorts

import "testing"

func QuckSort(left, right int, array *[6]int) {
	// 最左边数
	l := left
	// 最右边数
	r := right

	// 找到中间的数
	pivot := array[(left+right)/2]

	// 将pivot 小的数放到左边,大的数放到右边
	for l < r {
		// 从pivot左边找到大于等于值
		for array[l] < pivot {
			l++
		}
		// 从pivo右边找到小于等于值
		for array[r] > pivot {
			r--
		}
		// 表明本地分割任务完成
		if l >= r {
			break
		}
		// 交换
		array[l], array[r] = array[r], array[l]
		// 优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// l == r 移动一位
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuckSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuckSort(l, right, array)
	}
}
func TestQucikSort(t *testing.T) {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	t.Log(arr)
	QuckSort(0, len(arr)-1, &arr)
	t.Log(arr)
}
