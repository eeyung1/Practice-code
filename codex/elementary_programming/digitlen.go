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

	// negative := false
	var result []rune
	if n < 0 {
		//negative = true
		n = -n
	}

	for n > 0 {
		remainder := n % base
		var char rune
		n /= base
		if remainder < 10 {
			char = rune('0' + remainder)
		} else {
			char = rune('a' + (remainder - 10))
		}

		result = append(result, char)
		//negative = true
	}

	return string(result)

}

func main() {
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
