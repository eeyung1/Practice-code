package main

import (
	"fmt"
	"strings"
)

func bruteForceCaesar(ciphertext string) string {
	decrypted := []string{}
	for shift := 1; shift <= 26; shift++ {
		decrypted = append(decrypted, fmt.Sprintf("Shift %d: %s", shift, CaesarCipher(ciphertext, -shift)))
	}

	var result strings.Builder

	for _, validtext := range decrypted {
		result.WriteString(validtext)
		result.WriteString("\n")
		result.WriteString("\n")

	}

	return result.String()
} 