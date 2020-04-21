package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Take1(args []string) []string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return []string{s}
}

func Take2(args []string) (result []string) {
	result = append(result, args[0])
	result = append(result, strings.Join(args[1:], " "))
	return
}

func Take3(args []string) (result []string) {
	for index, value := range args {
		result = append(result, fmt.Sprintf("[%d] %s", index, value))
	}
	return
}

func main() {
	start := time.Now()
	for _, s := range Take1(os.Args) {
		fmt.Println(s)
	}
	for _, s := range Take2(os.Args) {
		fmt.Println(s)
	}
	for _, s := range Take3(os.Args) {
		fmt.Println(s)
	}
	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())
}
