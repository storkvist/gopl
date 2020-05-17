package main

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}

	if s[0] == '+' {
		s = s[1:]
	}

	fraction := ""
	if dot := strings.LastIndex(s, "."); dot > 0 {
		fraction = s[dot:]
		s = s[:dot]
	}

	counter := 3 - len(s)%3
	for i, c := range s {
		buf.WriteRune(c)
		counter++
		if i != len(s)-1 && counter%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String() + fraction
}
