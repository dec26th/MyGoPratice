package bench_connect_strings

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	result := ""
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result += "a"
	}
}

func BenchmarkSprintf(b *testing.B) {
	result := ""
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%s%s",result, "a")
	}
}

func BenchmarkBuffer(b *testing.B) {
	bb := bytes.Buffer{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bb.WriteString("a")
	}
}

func BenchmarkBuffer2(b *testing.B) {
	bb := bytes.Buffer{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bb.Write([]byte("a"))
	}
}
