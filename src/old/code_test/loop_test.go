package code_test

import "testing"

// 循环测试
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

//无限循环
func TestWhileLoop1(t *testing.T) {
	n := 0
	for {
		t.Log(n)
		n++
	}
}