package main

import (
	"fmt"
)

func ZipString(s string) string {
	current := s[0]
	count := 1
	result := ""

	for i := 1; i < len(s); i++ {
		char := s[i]
		if char == current {
			count++
		} else {
			result += string(rune(count + '0')) + string(current)

			current = char
			count = 1
		}
	}

	result += string(rune(count + '0')) + string(current)

	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}