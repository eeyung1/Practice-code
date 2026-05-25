package main

import (
	"fmt"
)

func main() {
	secretMessage := "hello world from golang"

	key := 3

	fmt.Println("Original Text:", secretMessage)

	encrypted := encryptCaesar(secretMessage, key)
	fmt.Println("Encrypted Text:", encrypted)

	decrypted := decryptCaesar(encrypted, key)
	fmt.Println("Decrypted Text:", decrypted)
}

func encryptCaesar(input string, shift int) string {
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] + rune(shift)

		if runes[i] > 'z' {
			runes[i] = runes[i] - 26
		}
	}

	return string(runes)
}

func decryptCaesar(input string, shift int) string {
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] - rune(shift)

		if runes[i] < 'a' {
			runes[i] = runes[i] + 26
		}
	}

	return string(runes)
}
