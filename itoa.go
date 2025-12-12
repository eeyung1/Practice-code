package main

import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		digit := n%10
		result = string('0' + digit) + result
		n /= 10
	}

	if isNegative == true {
		result = "-" + result
	}
	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

