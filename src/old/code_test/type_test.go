package code_test

import "testing"

// 按位清零
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a&^ Readable
	a = a&^ Writeable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	t.Log(Readable, Writeable, Executable)
}
