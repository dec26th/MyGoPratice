package string_to_byteslice_without_memalloc

import (
	"testing"
)

var testStc = "12323523rlkwnflsdjp23jlkqnejk123kl1 2ei1u2h1 2k312o31j2lwn12l3j12lk3nmenl1k2j3l12elk12lk3j1l2nwlmn12kjw1"

func BenchmarkS2B(b *testing.B) {
	result := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		result[i] = testStc
	}
	test := make([][]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range result {
		test[i] = s2b(result[i])
	}
}

func BenchmarkDirectS2B(b *testing.B) {
	ss := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		ss[i] = testStc
	}
	sb := make([][]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := range ss {
		sb[i] = []byte(ss[i])
	}
}

func BenchmarkB2S(b *testing.B) {
	sb := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		sb[i] = []byte(testStc)
	}
	ss := make([]string, b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := range sb {
		ss[i] = b2s(sb[i])
	}
}

func BenchmarkDirectB2S(b *testing.B) {
	sb := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		sb[i] = []byte(testStc)
	}
	ss := make([]string, b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := range sb {
		ss[i] = string(sb[i])
	}
}
