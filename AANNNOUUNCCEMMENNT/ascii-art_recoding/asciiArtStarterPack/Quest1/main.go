package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: please provide a word")
		return
	}

	word := os.Args[1]

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
}