package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	if s == "" {
		return ""
	}

	result := ""

	for _, c := range s {
		position := 1

		if c >= 'a' && c <= 'z' {
			position = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			position = int(c - 'A' + 1)
		}

		for i := 0; i < position; i++ {
			result += string(c)
		}
	}

	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
