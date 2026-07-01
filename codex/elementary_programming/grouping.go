package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	pattern := os.Args[1]
	text := os.Args[2]

	if !strings.HasPrefix(pattern, "(") || !strings.HasSuffix(pattern, ")") {
		return
	}

	if len(text) == 0 {
		return
	}

	inner := pattern[1:len(pattern)-1]
	terms := strings.Split(inner, "|")

	words := []string{}
	currentWord := ""

	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '\'' || ch == '’' {
			currentWord += string(ch)
		} else {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	counter := 1
	result := []string{}

	for _, word := range words {
		for _, term := range terms {
			if strings.Contains(word, term) {
				result = append(result, fmt.Sprintf("%d: %s", counter, word))
				counter++
			}
		}
	}

	if len(result) == 0 {
		return
	}

	for _, line := range result {
		fmt.Println(line)
	}
}