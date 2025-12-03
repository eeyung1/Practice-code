package main

import "fmt"

func FirstWord(s string) string {
	i := 0
	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[:i] + "\n"
}