package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	//ZiB
	//YiB
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	//ZB = 1000 * EB
	//YB = 1000 * ZB
)

func main() {
	fmt.Printf("KiB = %v\n", KiB)
	fmt.Printf("MiB = %v\n", MiB)
	fmt.Printf("GiB = %v\n", GiB)
	fmt.Printf("TiB = %v\n", TiB)
	fmt.Printf("PiB = %v\n", PiB)
	fmt.Printf("EiB = %v\n", EiB)
	//fmt.Printf("ZiB = %v\n", ZiB)
	//fmt.Printf("YiB = %v\n", YiB)

	fmt.Printf("KB = %v\n", KB)
	fmt.Printf("MB = %v\n", MB)
	fmt.Printf("GB = %v\n", GB)
	fmt.Printf("TB = %v\n", TB)
	fmt.Printf("PB = %v\n", PB)
	fmt.Printf("EB = %v\n", EB)
	//fmt.Printf("ZB = %v\n", ZB)
	//fmt.Printf("YB = %v\n", YB)
}
