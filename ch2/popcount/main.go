package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountByShifting calculates number of set bits using right shift.
func PopCountByShifting(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return
}

// PopCountByClearing calculates number of set bits with clearing of rightmost non-zero bit
func PopCountByClearing(x uint64) (count int) {
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return
}
