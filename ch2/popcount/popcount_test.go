package popcount_test

import (
	"testing"

	"github.com/storkvist/gopl/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClearing(0x1234567890ABCDEF)
	}
}
