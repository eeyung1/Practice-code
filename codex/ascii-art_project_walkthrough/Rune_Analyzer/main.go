package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No argument provided")
		return
	}

	input := os.Args[1]

	word := []rune(input)

	letter := 0
	num := 0
	punct := 0
	fmt.Println("Length:", len(word))
	for _, r := range word {
		if unicode.IsLetter(r) {
			letter++
		} else if unicode.IsDigit(r) {
			num++
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			punct++
		}
	}

	fmt.Println("Letters:", letter)
	fmt.Println("Digits:", num)
	fmt.Println("Symbols:", punct)
	fmt.Println("First char:", string(word[0]))
	fmt.Println("Last char:", string(word[len(word)-1]))
}
