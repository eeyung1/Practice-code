package main

import (
	"fmt"
	"os"

)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Error: No word provided")
		return
	}

	word := os.Args[1]


	counts := make(map[rune]int)

	order := []rune{}

	for _, char := range word {
		if counts[char] == 0 {
			order = append(order, char)
		}

		counts[char]++
	}

	// fmt.Println(order)

	for _, char := range order {
		fmt.Printf("%s -> %d\n", string(char), counts[char])
	}
}