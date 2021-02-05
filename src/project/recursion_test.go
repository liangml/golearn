package project

import (
	"fmt"
	"testing"
)

func recurlsion(n int) {
	if n > 2 {
		n--
		recurlsion(n)
	} else {
		fmt.Printf("n = %v\n", n)
	}
}

// 递归函数
func TestRecursion(t *testing.T) {
	recurlsion(4)
}

// 递归方式求出斐波那契数 1,1,2,3,5,8,13....

func Fib(n int) int {
	if (n == 1) || (n == 2) {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
func TestFib(t *testing.T) {
	t.Log(Fib(6))
}
