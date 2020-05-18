package main

func rotate(a []int, n int) {
	length := len(a)
	if length < 2 || n%length == 0 {
		return
	}
	if (n > 0 && n > length) || -n > length {
		n = n / length
	}
	if n < 0 {
		n = length + n
	}

	i, previous := length-1, a[length-1]
	for n > 0 {
		j := i - n
		if j < 0 {
			j = j + length
		}

		i, a[j], previous = j, previous, a[j]

		n--
	}
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func rotateWithReverse(a []int, n int) {
	length := len(a)
	if length < 2 || n%length == 0 {
		return
	}

	if (n > 0 && n > length) || -n > length {
		n = n / length
	}

	if n < 0 {
		n = length + n
	}

	reverse(a[:n])
	reverse(a[n:])
	reverse(a)
}
