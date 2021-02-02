package project

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// 自定义sort方法
type hero struct {
	Name string
	Age  int
}
type heroSlice []hero

func (hs heroSlice) Len() int {
	return len(hs)
}
func (hs heroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs heroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func TestInterFaces(t *testing.T) {
	arr := []int{24, 69, 80, 57, 13}

	// sort 冒泡排序
	sort.Ints(arr)
	t.Log(arr)

	var heros heroSlice
	for i := 0; i < 10; i++ {
		hero := hero{
			Name: fmt.Sprintf("英雄｜%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	// 排序前
	fmt.Println("排序前")
	for _, v := range heros {
		fmt.Println(v)
	}
	sort.Sort(heros)
	// 排序后
	fmt.Println("排序后")
	for _, v := range heros {
		fmt.Println(v)
	}
}
