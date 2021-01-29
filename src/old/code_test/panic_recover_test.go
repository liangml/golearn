package code_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong"))
	//os.Exit(-1)
}
