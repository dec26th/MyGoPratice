package bench_connect_strings

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	result := ""
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result += strconv.Itoa(i)
	}
}

func BenchmarkSprintf(b *testing.B) {
	result := ""
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%s%s",result, strconv.Itoa(i))
	}
}

func BenchmarkBuffer(b *testing.B) {
	bb := bytes.Buffer{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bb.WriteString(strconv.Itoa(i))
	}
}

func BenchmarkBuffer2(b *testing.B) {
	bb := bytes.Buffer{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bb.Write([]byte(strconv.Itoa(i)))
	}
}
