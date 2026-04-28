package main

import (
	// "fmt"
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	output := make([]string, 8)

	if input == "" {
		return output
	}

	for i := 0; i < 8; i++ {
		var builder strings.Builder

		for _, ch := range input {
			art, ok := banner[ch]

			if !ok {
				art = banner[' ']
			}

			builder.WriteString(art[i])
		}

		output[i] = builder.String()
	}

	return output
}


// func RenderLine(input string, banner map[rune][]string) []string {
// 	output := make([]string, 8)

// 	for _, ch := range input {
// 		art, ok := banner[ch]
// 		if !ok {
// 			art = banner[' ']
// 		}

// 		if art == nil {
// 			continue
// 		}

// 		for i := range 8 {
// 			output[i] += art[i]
// 		}
// 	}

// 	return output

// }

// func Renderlines(input string, banner map[rune][]string) []string {
// 	output := make([]string, 8)

// 	for _, ch := range input {
// 		art, ok := banner[ch]

// 		if !ok {
// 			art = banner[' ']
// 		}

// 		if art == nil {
// 			continue
// 		}

// 		for i := range 8 {
// 			output[i] += art[i]
// 		}
// 	}

// 	return output
// }

// func Renderlines(input string, banner map[rune][]string) []string {
// 	output := make([]string, 8)

// 	for _, ch := range input {
// 		//art := banner[ch]

// 		for i := range 8 {
// 			// output[i] += art[i]
// 			output[i] += banner[ch][i]
// 		}
// 	}

// 	return output
// }
