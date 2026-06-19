package main

import (
	"fmt"
)

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output"
	}

	result := []string{}

	currentWord := ""

	for _, ch := range str {
		if ch == ' ' {
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}

	if currentWord != "" {
		result = append(result, currentWord)
	}

	final := ""

	for i := len(result) - 1; i >= 0; i-- {
		if i == len(result)-1 {
			final += result[i]
		} else {

			final += " " + result[i]
		}
	}

	return final + "\n"
}

func main() {
	fmt.Print(WordFlip("First            second last"))
	fmt.Println(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
