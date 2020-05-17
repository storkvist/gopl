package main

func areAnagram(a, b string) bool {
	length := len(a)
	if length != len(b) {
		return false
	}

	chars := make(map[rune]int)
	for _, c := range a {
		chars[c]++
	}
	for _, c := range b {
		chars[c]--

		if n, _ := chars[c]; n == 0 {
			delete(chars, c)
		}
	}

	return len(chars) == 0
}
