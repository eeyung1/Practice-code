package main

import (
	"fmt"
)

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	negative := false
	if value < 0 {
		negative = true
		value = -value
	}

	digits := "0123456789ABCDEF"
	result := []rune{}

	for value > 0 {
		remainder := value % base
		result = append(result, rune(digits[remainder]))
		value = value / base
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Add minus sign if negative
	if negative {
		return "-" + string(result)
	}

	return string(result)
}

func main() {
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-42, 4))
	fmt.Println(ItoaBase(123, 10))
	fmt.Println(ItoaBase(0, 8))
	fmt.Println(ItoaBase(255, 2))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(15, 16))
	fmt.Println(ItoaBase(10, 4))
	fmt.Println(ItoaBase(255, 10))
}
