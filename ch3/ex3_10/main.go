package main

import "bytes"

func comma(s string) string {
	var buf bytes.Buffer
	counter := 3 - len(s)%3
	for i, c := range s {
		buf.WriteRune(c)
		counter++
		if i != len(s)-1 && counter%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
