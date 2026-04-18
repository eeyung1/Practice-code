package main

import (
	"fmt"
)

func main() {
	var num1 []rune
	var num2 []rune
	for i := 1; i <= 22; i += 1 {

		if i%2 != 0 {
			num1 = append(num1, rune(i))
		} else {
			num2 = append(num2, rune(i))
		}
	}

	fmt.Println(num1)
	fmt.Println(num2)

	var result []rune

	maxLen := len(num1)
	if len(num2) > maxLen {
		maxLen = len(num2)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(num1) && i < len(num2) {
			result = append(result, num1[i]*num2[i])
		} else if i < len(num1) {
			result = append(result, num1[i])
		} else if i < len(num2) {
			result = append(result, num2[i])
		}
	}

	fmt.Println(result)
}
