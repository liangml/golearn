package unit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func square(op int) int {
	return op*op + 1
}

func TestSquare(t *testing.T) {
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d ,the actual %d", inputs[i], expected[i], ret)
		}
	}
}
func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("end")
}
func TestFailInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("Error")
	fmt.Println("end")
}

func TestSquareWithAssert(t *testing.T) {
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
