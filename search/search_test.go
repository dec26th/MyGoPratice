package search

import (
	"testing"
)

func BenchmarkSearchByDoubleLoop(b *testing.B) {
	target := make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		target[i] = byte(1)
	}

	bench := make([]byte, b.N)
	copy(bench, target)
	filter := getDoubleLoop(target)
	filter(bench)
}

func BenchmarkSearchByMap(b *testing.B) {
	target := make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		target[i] = byte(1)
	}

	bench := make([]byte, b.N)
	copy(bench, target)
	filter := getMap(target)
	filter(bench)
}
