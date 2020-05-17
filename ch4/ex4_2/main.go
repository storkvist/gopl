package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	SHA256 = 256
	SHA384 = 384
	SHA512 = 512
)

var algorithm int

func init() {
	flag.IntVar(&algorithm, "sha", SHA256, "Which SHA function should be used (256, 384 or 512)")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		switch algorithm {
		case SHA256:
			fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(s)))
		case SHA384:
			fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(s)))
		case SHA512:
			fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(s)))
		default:
			log.Fatal("Unknown SHA function")
		}
	}
}
