package project_test

import (
	"testing"
)

// 每一位数都是前一位数字的和
// 1,1,2,3,5,8,13

func TestFibList(t *testing.T) {
	var (
		a int = 1
	)
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 1
	a, b = b, a
	t.Log(a, b)
}
