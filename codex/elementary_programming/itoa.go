package main

import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := n < 0

	if n < 0 {
		n = -n
	}

	result := ""
	final := ""

	for n > 0 {
		digit := n % 10
		result += string(rune(digit + '0'))
		n /= 10
	}

	for i := len(result)-1; i >= 0; i-- {
		final += string(result[i])
	}

	if isNegative {
		final = "-" + final
	}

	return final
}


func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

// func Itoa(n int) string {
// 	result := fmt.Sprintf("%d", n)
// 	return result
// }