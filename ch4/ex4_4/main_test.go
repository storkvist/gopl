package main

import "testing"

var tests = []struct {
	a    []int
	n    int
	want []int
}{
	{[]int{}, 1, []int{}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
}

func TestRotate(t *testing.T) {
	for _, test := range tests {
		a := make([]int, len(test.a))
		copy(a, test.a)
		rotate(a, test.n)
		if len(a) != len(test.a) {
			t.Errorf("rotate()")
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate([]int{1, 2, 3, 4, 5}, 2)
	}
}

func BenchmarkRotateWithReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateWithReverse([]int{1, 2, 3, 4, 5}, 2)
	}
}
