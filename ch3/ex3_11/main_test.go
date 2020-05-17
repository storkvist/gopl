package main

import "testing"

var tests = []struct {
	s    string
	want string
}{
	{"12345.456", "12,345.456"},
	{"-12345.456", "-12,345.456"},
	{"+12345.456", "12,345.456"},
}

func TestComma(t *testing.T) {
	for _, test := range tests {
		if got := comma(test.s); got != test.want {
			t.Errorf("comma(%v) = %v; want %v", test.s, got, test.want)
		}
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma("-12345.456")
	}
}
