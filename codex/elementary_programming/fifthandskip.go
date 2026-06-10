package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Input\n"
	}

	value := []rune{}
	result := ""

	for _, ch := range str {
		if ch != ' ' {
			value = append(value, ch)
		}
	}

	keptCount := 0



	for i, ch := range value {
		if i %6 == 5 {
			continue
		}

		result += string(ch)
		keptCount++

		if keptCount%5 == 0 && i < len(value)-1 {
			result += " "
		}
	}





	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
