package test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStrconvItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}
}

func BenchmarkStrconvFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkStrconvFormatUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatUint(uint64(i), 10)
	}
}

func BenchmarkSprinf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}
