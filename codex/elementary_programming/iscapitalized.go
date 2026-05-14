package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	for i, r := range s {
		isLower := r >= 'a' && r <= 'z'

		if i == 0 && isLower {
			return false
		}
		if i > 0 && s[i-1] == ' ' && isLower {
			return false
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
