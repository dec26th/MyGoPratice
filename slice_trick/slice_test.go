package slice_trick

import (
	"fmt"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testSlice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := 4

	testSlice1 := append(testSlice[:i], testSlice[i+1:]...)
	testSlice3 = append(testSlice3[:i], testSlice3[i+1:]...)
	fmt.Println(testSlice)
	fmt.Println(testSlice1)
	fmt.Println(testSlice3)
	fmt.Println(cap(testSlice1), len(testSlice1))
	fmt.Println(cap(testSlice3), len(testSlice3))
}

func TestSliceDelete2(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	i := 4

	copy(testSlice[i:], testSlice[i+1:])
	fmt.Println(testSlice)
}

func TestSliceCopy(t *testing.T) {
	testSLice := []int{1, 2, 3, 4, 5, 6, 7}
	test1 := testSLice[:3]
	fmt.Println(len(test1), cap(test1))
	test1[0] = 11111
	fmt.Println(testSLice)
}
