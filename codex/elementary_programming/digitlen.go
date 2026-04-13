package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	count := 0
	for n > 0 {
		n /= base
		count++
	}
	return count
}

func IntToBase(n, base int) string {
	if base < 2 || base > 36 {
		return ""
	}

	if n == 0 {
		return "0"
	}

	negative := false
	var result []rune

	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		remainder := n % base
		var char rune

		if remainder < 10 {
			char = rune('0' + remainder)
		} else {
			char = rune('a' + (remainder - 10))
		}

		result = append(result, char)
		n /= base
	}

	// reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// handle negative
	if negative {
		result = append([]rune{'-'}, result...)
	}

	return string(result)
}

func BaseToInt(s string, base int) int {
	if base < 2 || base > 36 {
		return 0
	}

	if s == "" {
		return 0
	}

	result := 0
	start := 0
	negative := false

	if s[0] == '-' {
		negative = true
		start = 1
	}

	// handle case like "-"
	if start == len(s) {
		return 0
	}

	for i := start; i < len(s); i++ {
		r := rune(s[i])
		var digit int

		if r >= '0' && r <= '9' {
			digit = int(r - '0')
		} else if r >= 'a' && r <= 'z' {
			digit = int(r-'a') + 10
		} else if r >= 'A' && r <= 'Z' {
			digit = int(r-'A') + 10
		} else {
			return 0
		}

		if digit >= base {
			return 0
		}

		result = result*base + digit
	}

	if negative {
		result = -result
	}

	return result
}

func main() {
	fmt.Println(BaseToInt("1100100", 2))
	fmt.Println(BaseToInt("ff", 16))
	fmt.Println(BaseToInt("-42", 10))
	fmt.Println(IntToBase(100, 10))
	fmt.Println(IntToBase(100, 2))
	fmt.Println(IntToBase(255, 16))
	fmt.Println(IntToBase(-42, 10))
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
	fmt.Println(DigitLen(0, 10))
	fmt.Println(DigitLen(0, 0))
}
