package main

import (
	"fmt"
)


func PrintMemory(arr [10]byte) {

	for i, r := range arr {
		digit1 := r / 16
		digit2 := r % 16

		var hex1 rune
		var hex2 rune

		if digit1 < 10 {
			hex1 = rune('0' + digit1)
		} else {
			hex1 = rune('a' + (digit1 - 10))
		}

		if digit2 < 10 {
			hex2 = rune('0' + digit2)
		} else {
			hex2 = rune('a' + (digit2 - 10))
		}

		fmt.Print(string(hex1) + string(hex2))

		if (i+1)%4 != 0 && i != len(arr) - 1 {
			fmt.Print(" ")
		}


		if (i+1)%4 == 0 || i == len(arr)-1 {
			fmt.Println()
		}
	}


	for _, r := range arr {
		if r >= 32 && r <= 126 {
			fmt.Print(string(r))
		} else {
			fmt.Print(".")
		}
	}

	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
