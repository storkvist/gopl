package main

import "testing"

var tests = []struct {
	a    [5]int
	want [5]int
}{
	{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
	{[5]int{2, 3, 4, 5, 1}, [5]int{1, 5, 4, 3, 2}},
}

func TestReverse(t *testing.T) {
	for _, test := range tests {
		a := test.a
		reverse(&a)
		if a != test.want {
			t.Errorf("Reverse(%v) -> %v; want %v", test.a, a, test.want)
		}
	}
}
