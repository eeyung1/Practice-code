package main

import (
	"fmt"
	"os"
	"unicode"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("No argument provided")
		return
	}

	input := os.Args[1]

	runes := []rune(input)

	fmt.Println("Length:", len(runes))

	count := 0
	num := 0

	for _, r := range runes {
		if unicode.IsLetter(r) {
			count++
		}

		if unicode.IsDigit(r) {
			num++
		}
	}

	fmt.Println("Letters:", count)
	fmt.Println("Digits:", num)
	fmt.Println("First char:", string(runes[0]))
	fmt.Println("Last char:", string(runes[len(runes)-1]))
}


/*

NOTE: 

When working with characters, it always necessary to convert
the characters to runes(unicode characters) so it can handle
a wider range of characters like emojis and others.



*/