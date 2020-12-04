package _type

import (
	"fmt"
	"testing"
)

type a = int
type b int

func TestType(t *testing.T) {
	test := a(1)
	test2 := b(1)
	test1 := 1
	fmt.Println(test == test1)
	fmt.Println(test2 == test)
}

func TestPtr(t *testing.T) {
	var tPtr testPtr = &test{}
	tPtr.T()
	fmt.Println(tPtr)
}
