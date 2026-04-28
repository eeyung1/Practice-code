package main

// import (
// 	"fmt"
// 	"strings"
// )

func Renderlines(input string, banner map[rune][]string)[]string {
	output := make([]string, 8)

	for _, ch := range input {
		art := banner[ch]
		for i := 0; i < 8; i++ {
			output[i] += art[i]
		}
	}

	return output
}