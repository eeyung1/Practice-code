package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	result := ""
	firstWord := true

	i := 0
	for i < len(input) {
		for i < len(input) && isSpace(rune(input[i])) {
			i++
		}

		start := i

		for i < len(input) && !isSpace(rune(input[i])) {
			i++
		}

		if start < i {
			if !firstWord {
				result += "   "
			}
			result += input[start:i]
			firstWord = false
		}
	}

	fmt.Println(result)
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r'
}
