package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]

	if input == "" {
		fmt.Println()
		return
	}

	var result []rune
	inWord := false

	for _, ch := range input {
		if isSpace(ch) {
			if inWord {
				result = append(result, ' ')
				inWord = false
			}
		} else {
			result = append(result, ch)
			inWord = true
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)]
	}

	fmt.Println(string(result))
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r'
}
