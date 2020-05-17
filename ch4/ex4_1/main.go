package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func setBitsCount(x byte) int {
	return int(pc[x])
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", c1, c2)

	count := 0
	for i, b1 := range c1 {
		count += setBitsCount(b1 ^ c2[i])
	}
	fmt.Printf("count of different bits = %v\n", count)
}
