package unit

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSquare(t *testing.T) {
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d ,the expected is %d,the actual %d", inputs[i], expected[i], ret)
		}
	}
}
func TestErrorIncode(t *testing.T) {
	fmt.Print("start")
	t.Error("Error")
	fmt.Println("end")
}
func TestFailIncode(t *testing.T) {
	fmt.Print("start")
	t.Fatal("Error")
	fmt.Println("end")
}
func TestSquareWithAssert(t *testing.T) {
	// 表格测试法
	// 使用公共库测试
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
