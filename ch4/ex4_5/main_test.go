package main

import (
	"testing"
)

var tests = []struct {
	a    []string
	want []string
}{
	{[]string{"1", "1", "2", "3", "3", "2", "3"}, []string{"1", "2", "3", "2", "3"}},
}

func TestSquezze(t *testing.T) {
	for _, test := range tests {
		a := make([]string, len(test.a))
		copy(a, test.a)
		a = squezze(a)
		for i := range test.want {
			if a[i] != test.want[i] {
				t.Errorf("squezze(%v) -> %v; want %v\n", test.a, a, test.want)
				continue
			}
		}
	}
}
