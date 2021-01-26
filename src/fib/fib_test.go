package fib

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// fibonacci数列 斐布拉切数列实现
func TestFilbList(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Println("", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	// tmp := a
	// a = b
	// b = tmp
	t.Log(a, b)
}

// iota位运算
const (
	Monday = iota + 1
	TuesDay
	Wednesday
)
const (
	Readale = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, TuesDay, Wednesday)
}
func TestConstantTry1(t *testing.T) {
	a := 1 //7
	t.Log(a&Readale == Readale, a&Writable == Writable, a&Executable == Executable)
}

// 按位置清零 &^
func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readale
	a = a &^ Writable
	t.Log(a&Readale == Readale, a&Writable == Writable, a&Executable == Executable)
}

// map实现工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// 使用map实现set功能 map[type]bool
// 元素唯一性，添加元素、判断元素是否存在、删除元素、元素个数
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[1] {
		t.Logf("%d is exsits", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[1] {
		t.Logf("%d is exsits", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

// 计时函数实现类似装饰着模式
// 自定义类型
type ConvInt func(op int) int

func timeSpent(inner ConvInt) ConvInt {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("运行时长为:", time.Since(start))
		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParm(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}
func Clear() {
	fmt.Println("Clear resources.")
}

// 延迟执行defer
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}

// interface 多态
// interface 对应只能是指针
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct{}
type JavaProgrammer struct{}

func (c *GoProgrammer) WriteHelloWorld() Code {
	return "go hello world"
}
func (c *JavaProgrammer) WriteHelloWorld() Code {
	return "java hello world"
}
func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestClient(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}

// 扩展和复用  继承
type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}
func (p *Pet) Speakto(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speakto("liangml")
}

// 空接口断言
func DoSomething(p interface{}) {
	// 第一种方法
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("unknow Type")

	// 第二种方法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow Type")
	}
}
func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

// 接口最佳实践 最小接口
// 使用小的接口定义，只包含一个方法
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 较大接口定义，由多个小接口组合而成
type ReadWriter interface {
	Reader
	Writer
}

// 只依赖于必要功能的最小接口
func StoreData(reader Reader) error {
	return nil
}

// 错误检查机制
var (
	// LessThanTwoError 预制错误信息
	ErrLessThanTwo = errors.New("n should be not less than 2")
	// LargerThenHundredError 预制错误信息
	ErrLargerThenHundred = errors.New("n should be not larger than 100")
)

func GetFibonacci(n int) ([]int, error) {
	// 及早失败，避免嵌套
	if n < 2 {
		return nil, ErrLessThanTwo
	}
	if n > 100 {
		return nil, ErrLargerThenHundred
	}
	filbList := []int{1, 1}
	for i := 2; /*短变量声明*/ i < n; i++ {
		filbList = append(filbList, filbList[i-2]+filbList[i-1])
	}
	return filbList, nil
}
func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-5); err != nil {
		if err == ErrLessThanTwo {
			fmt.Println("it is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// panic recover
func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something Wrong"))
}
