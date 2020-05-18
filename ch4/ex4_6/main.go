package main

import (
	"unicode"
	"unicode/utf8"
)

func compressUnicodeSpaces(s []byte) []byte {
	length := len(s)
	startOfSpaces := -1
	for i := 0; i < length; {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			if startOfSpaces < 0 {
				startOfSpaces = i
			}
		} else {
			if startOfSpaces >= 0 {
				s[startOfSpaces] = ' '
				length++
				copy(s[startOfSpaces+1:], s[i:])
				length -= i - startOfSpaces
				startOfSpaces = -1
			}
		}

		i += size
	}

	if startOfSpaces >= 0 {
		s[startOfSpaces] = ' '
		length = startOfSpaces + 1
	}

	return s[:length]
}
