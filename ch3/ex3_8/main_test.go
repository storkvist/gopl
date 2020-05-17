package main

import "testing"

func BenchmarkImg64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img64()
	}
}

func BenchmarkImg128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img128()
	}
}
