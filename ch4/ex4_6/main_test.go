package main

import "testing"

var tests = []struct {
	s    string
	want string
}{
	{"Hello," + string('\u0085') + string('\u00a0') + string('\u1680') + "World!", "Hello, World!"},
	{string('\u0085') + "Hello", " Hello"},
	{"Hello" + string('\u0085') + string('\u1680'), "Hello "},
}

func TestCompressUnicodeSpaces(t *testing.T) {
	for _, test := range tests {
		s := []byte(test.s)
		got := compressUnicodeSpaces(s)
		want := []byte(test.want)
		if len(got) != len(want) {
			t.Errorf("compressUnicodeSpaces(%q) = %q; want %q", test.s, string(got), test.want)
			continue
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("compressUnicodeSpaces(%q) = %q; want %q", test.s, string(got), test.want)
				continue
			}
		}
	}
}
