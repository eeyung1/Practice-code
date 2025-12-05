package main

import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var bytes []byte
	isNegative := n < 0

	if n > 0 {
		n = -n
	}

	for n != 0 {
		digit := n % 10
		bytes = append(bytes, byte('0'-digit))
		n /= 10
	}

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	if isNegative {
		bytes = append([]byte{'-'}, bytes...)
	}

	return string(bytes)
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
