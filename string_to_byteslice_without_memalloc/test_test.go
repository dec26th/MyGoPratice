package string_to_byteslice_without_memalloc

import (
	"fmt"
	"strconv"
	"testing"
)

const num = 1000000

func TestS2B(t *testing.T) {
	s := "123"
	fmt.Println(s2b(s))
}

func BenchmarkS2B(b *testing.B) {
	result := make([]string, num)
	for i := 0; i < num; i++ {
		result[i] = strconv.Itoa(i)
	}
	test := make([][]byte, num)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range result {
		test[i] = s2b(result[i])
	}
}

func BenchmarkDirect(b *testing.B) {
	result := make([]string, num)
	for i := 0; i < num; i++ {
		result[i] = strconv.Itoa(i)
	}
	test := make([][]byte, num)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range result {
		test[i] = []byte(result[i])
	}
}