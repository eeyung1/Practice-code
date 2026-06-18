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

	words := []string{}
	currentWord := ""

	for _, ch := range input {
		if ch == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		fmt.Println()
		return
	}

	for i := 1; i < len(words); i++ {
		fmt.Print(words[i], " ")
	}


	fmt.Print(words[0])
	fmt.Println()
}