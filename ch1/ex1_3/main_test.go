package main

import "testing"

var args = []string{"program", "par1", "par2", "par3"}

func BenchmarkTake1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Take1(args)
	}
}

func BenchmarkTake2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Take2(args)
	}
}

func BenchmarkTake3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Take3(args)
	}
}
