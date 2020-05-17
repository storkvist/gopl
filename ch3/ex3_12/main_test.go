package main

import "testing"

var tests = []struct {
	a    string
	b    string
	want bool
}{
	{"abc", "bca", true},
	{"abc", "abd", false},
	{"👨‍👩‍👦ab", "a👨‍👩‍👦b", true},
}

func TestAreAnagramm(t *testing.T) {
	for _, test := range tests {
		if got := areAnagram(test.a, test.b); got != test.want {
			t.Errorf("areAnagram(%v, %v) = %v; want %v\n", test.a, test.b, got, test.want)
		}
	}
}

func BenchmarkAreAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areAnagram("👨‍👩‍👦ab", "a👨‍👩‍👦b")
	}
}
