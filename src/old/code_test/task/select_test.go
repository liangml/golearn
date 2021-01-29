package task_test

import (
	"fmt"
	"testing"
	"time"
)

func service1() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

// 异步调用
func AsyncService1() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service1()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}


// 利用select实现多路调用超时
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService1():
		t.Log(ret)
	case <-time.After(time.Microsecond * 100):
		t.Error("time out")
	}

}
