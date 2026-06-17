package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	words := []string{}
	currentWord := ""
	
	for _, ch := range input {
		if ch == ' ' {
			words = append(words, currentWord)
			currentWord = ""
		} else {
			currentWord += string(ch)
		}
	}

	fmt.Println(currentWord)

	words = append(words, currentWord)


	fmt.Print(words[len(words)-1]) // first word in reversed order

	for i := len(words) - 2; i >= 0; i-- {
		fmt.Print(" ", words[i])
	}
	fmt.Println()
}
