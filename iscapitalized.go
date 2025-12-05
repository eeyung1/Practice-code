package main

import (
	"fmt"
)

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	newWord := true

	for i := 0; i < len(s); i++ {
		c := s[i]

		if isAlphaNumeric(c) {
			if newWord {
				if isLower(c) {
					return false
				}
				newWord = false
			}
		} else {
			newWord = true
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
