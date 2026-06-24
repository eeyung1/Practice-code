package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]

	vowelPos := -1

	for i, ch := range word {
		lower := ch
		if ch >= 'A' && ch <= 'Z' {
			lower = ch + 32
		}

		if lower == 'a' || lower == 'e' || lower == 'i' || lower == 'o' || lower == 'u' {
			vowelPos = i
			break
		}
	}

	if vowelPos == -1 {
		fmt.Println("No vowels")
		return
	}

	var result string

	if vowelPos == 0 {
		result = word + "ay"
	} else {
		prefix := word[:vowelPos]
		suffix := word[vowelPos:]

		result = suffix + prefix + "ay"
	}

	fmt.Println(result)
}
