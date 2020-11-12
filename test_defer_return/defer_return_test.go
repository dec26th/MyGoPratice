package test_defer_return

import (
	"fmt"
	"testing"
)

func DeferAndReturn() int {
	return 1
	defer Defer()
	return 3
}

func Defer() int {
	return 2
}

func TestDeferAndReturn(t *testing.T) {
	fmt.Println(DeferAndReturn())
}
