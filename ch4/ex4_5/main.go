package main

func squezze(a []string) []string {
	length := len(a)
	if length == 0 {
		return a
	}

	i, j := 0, 0
	prev := a[0]
	for j < length {
		if a[j] != prev {
			a[i] = prev
			prev = a[j]
			i++
		}

		j++
	}

	return a[:i+1]
}
