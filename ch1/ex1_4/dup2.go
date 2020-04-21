package main

import (
	"bufio"
	"fmt"
	"os"
)

type locations map[string]int

func main() {
	counts := make(map[string]locations)
	// counts := make(map[string]int)
	// locations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, locations := range counts {
		if len(locations) > 0 {
			fmt.Printf(">>> %s\n", line)
			for filename, n := range locations {
				fmt.Printf("%d\t%s\n", n, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]locations) {
	input := bufio.NewScanner(f)
	var text string
	for input.Scan() {
		text = input.Text()
		if counts[text] == nil {
			counts[text] = make(locations)
		}
		counts[text][f.Name()]++
	}
	// NOTE: Ignoring potential errors from input.Err()
}
