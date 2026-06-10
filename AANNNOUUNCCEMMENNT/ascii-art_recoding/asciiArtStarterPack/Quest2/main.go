package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: please provide a word")
		return
	}

	word := os.Args[1]

	if !strings.Contains(word, ".txt") {

		counts := make(map[rune]int)

		order := []rune{}

		for _, ch := range word {
			if counts[ch] == 0 {
				order = append(order, ch)
			}

			counts[ch]++
		}

		for _, ch := range order {
			fmt.Printf("%s -> %d\n", string(ch), counts[ch])
		}
	} else {
		data, err := os.ReadFile(word)
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
}
