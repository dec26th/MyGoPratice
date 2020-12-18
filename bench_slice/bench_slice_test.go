package bench_slice

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	test := make([]int, b.N)

	for i := 0; i < len(test); i++ {
		test[i] = 1
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	test1 := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		test1 = append(test1, 1)
	}
}
