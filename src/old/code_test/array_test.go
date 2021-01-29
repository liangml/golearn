package code_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}

	//多维数组的声明
	b := [3]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(b, c)
}
func TestArrayCut(t *testing.T)  {
	a := [...]int{1,2,3,4,5}
	t.Log(a[1:2])  // 2
	t.Log(a[1:3])  // 2,3
	t.Log(a[1:len(a)])  // 2,3,4,5
	t.Log(a[1:])  // 2,3,4,5
	t.Log(a[:3])  // 1,2,3
	t.Log(a[:])  // 1,2,3,4,5
}
