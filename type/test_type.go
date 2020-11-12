package _type

type test struct {
	a int
	b int
}

type testPtr = *test

func (t testPtr) T() {
	t.b  =1
}