package code_test

import (
	"fmt"
	"testing"
)

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

func TestDog(t *testing.T)  {
	dog:= new(Dog)
	dog.Speak()
}