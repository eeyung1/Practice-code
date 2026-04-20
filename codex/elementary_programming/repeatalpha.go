package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	if s == "" {
		return ""
	}

	var result []rune 

	for _, r := range s {
		repeat := 1

		if r >= 'a' && r <= 'z' {
			repeat = int(r-'a') + 1
		} else if r >= 'A' && r <= 'Z' {
			repeat = int(r-'A') + 1
		}

		for i := 0; i < repeat; i++ {
			result = append(result, r)
		}
	}

	return string(result)	
}

func RepeatDigit(s string) string {
	if s == "" {
		return ""
	}

	var result []rune

	for _, r := range s {
		if r < '0' && r > '9' {
			continue
		}

		repeat := int(r - '0')

		for i := 0; i < repeat; i++ {
			result = append(result, r)
		}
	}

	return string(result) 
}

func SumDigits(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	for _, r := range s {
		num := int(r-'0')
		result += num
	}


	return result
}

func main() {
	fmt.Println(SumDigits("203"))
	fmt.Println(SumDigits("999"))
	fmt.Println(RepeatDigit("203"))
	fmt.Println(RepeatDigit("293"))
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}