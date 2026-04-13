package main

import (
	"fmt"
)

// func FirstWord(s string) string {
// 	if s == "" {
// 		return "\n"
// 	}

// 	word := strings.Fields(s)

// 	return word[0] + "\n"
// }

func FirstWord(s string) string {
	i := 0

	//skip leading spaces
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	start := i

	//collect word
	for i < len(s) && (s[i] != ' ' && s[i] != '\t') {
		i++
	}

	if start == i {
		return "\n"
	}

	return s[start:i] + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
	fmt.Print(FirstWord("   hello\tworld"))
}
