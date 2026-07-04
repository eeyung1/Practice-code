package main

import (
	"fmt"
	"unicode"
)

func AlphaCount(s string) int {
	count := 0

	for _, ch := range s {
		lower := unicode.ToLower(ch)

		if lower >= 'a' && lower <= 'z' {
			count++
		}
	}

	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}