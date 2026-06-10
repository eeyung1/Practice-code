package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File not provided")
		return
	}

	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	word := string(data)

	var cleanedStr strings.Builder

	for _, ch := range word {
		if ch != ' ' {
			cleanedStr.WriteString(string(ch))
		}
	}

	cleaned := cleanedStr.String()

	cleaned = strings.ToLower(cleaned)

	counts := make(map[rune]int)

	order := []rune{}

	for _, ch := range cleaned {
		if counts[ch] == 0 {
			order = append(order, ch)
		}

		counts[ch]++
	}

	for _, ch := range order {
		fmt.Printf("%s -> %d\n", string(ch), counts[ch])
	}

	
}