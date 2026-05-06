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
		// Skip spaces/tabs
		for i < len(input) && isSpace(rune(input[i])) {
			i++
		}

		start := i

		// Capture word
		for i < len(input) && !isSpace(rune(input[i])) {
			i++
		}

		// If we found a word
		if start < i {
			if !firstWord {
				result += "   "
			}
			result += input[start:i]
			firstWord = false
		}
	}

	// ONLY print if we found at least one word
	if result != "" {
		fmt.Println(result)
	}
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r'
}
