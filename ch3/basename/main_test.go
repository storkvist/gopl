package main

import "testing"

var tests = []struct {
	s    string
	want string
}{
	{"a", "a"},
	{"a.go", "a"},
	{"a/b/c.go", "c"},
	{"a/b.c.go", "b.c"},
}

func TestBasename(t *testing.T) {
	for _, test := range tests {
		if got := basename(test.s); got != test.want {
			t.Errorf("basename(%v) = %v; want %v", test.s, got, test.want)
		}
	}
}

func TestBasenameSimple(t *testing.T) {
	for _, test := range tests {
		if got := basenameSimple(test.s); got != test.want {
			t.Errorf("basenameSimple(%v) = %v; want %v", test.s, got, test.want)
		}
	}
}

var benchmarkArg = "a/b.c.go"

func BenchmarkBasename(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basename(benchmarkArg)
	}
}

func BenchmarkBasenameSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basenameSimple(benchmarkArg)
	}
}
