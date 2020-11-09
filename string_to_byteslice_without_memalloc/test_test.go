package string_to_byteslice_without_memalloc

import (
	"strconv"
	"testing"
)

// 275656681	         4.18 ns/op	       0 B/op	       0 allocs/op
func BenchmarkS2B(b *testing.B) {
	result := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		result[i] = strconv.Itoa(i)
	}
	test := make([][]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range result {
		test[i] = s2b(result[i])
	}
}

//BenchmarkDirect-12    	46649576	        24.5 ns/op	       8 B/op	       1 allocs/op
func BenchmarkDirectS2B(b *testing.B) {
	ss := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		ss[i] = strconv.Itoa(i)
	}
	sb := make([][]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range ss {
		sb[i] = []byte(ss[i])
	}
}

//479128580	        11.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkB2S(b *testing.B) {
	sb := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		sb[i] = []byte{byte(i)}
	}
	ss := make([]string, b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := range sb {
		ss[i] = b2s(sb[i])
	}
}

//100000000	        10.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDirectB2S(b *testing.B) {
	sb := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		sb[i] = []byte{byte(i)}
	}
	ss := make([]string, b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := range sb {
		ss[i] = string(sb[i])
	}
}