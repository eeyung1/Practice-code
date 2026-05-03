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
		result = result[:len(result)-1]
	}

	fmt.Println(string(result))
	fmt.Println(countWords(input))
	fmt.Println(extractWords(input))
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r'
}

func countWords(s string) int {
	if s == "" {
		return 0
	}

	inWord := false
	count := 0

	for _, r := range s {
		if isSpace(r) {
			if inWord {
				inWord = false
				continue
			}
		} else {
			if !inWord {
				count++
				inWord = true
			}
		}
	}

	return count
}

func extractWords(s string) []string {
	if s == "" {
		return nil
	}

	var words []string
	var current []rune
	inWord := false

	for _, ch := range s {
		if !isSpace(ch) {
			current = append(current, ch)
			inWord = true
		} else {
			if inWord {
				words = append(words, string(current))
				current = []rune{}
				inWord = false
			}
		}
	}

	if inWord {
		words = append(words, string(current))
	}

	return words
}
