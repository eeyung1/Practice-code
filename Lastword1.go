package main

import "fmt"

func LastWord(s string) string {
	word := ""

	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		word = string(s[i]) + word
		i--
	}
	return word + "\n"
}