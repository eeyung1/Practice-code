package main

import (
	"fmt"
	"strconv"
	"strings"
)

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func Capitalize(word string) string {

	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func extractNumber(token string) int {

	token = strings.Trim(token, "()")
	parts := strings.Split(token, ",")

	if len(parts) != 2 {
		return 1
	}

	n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 1
	}

	return n
}

func HandleConversions(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		// HEX conversion
		if words[i] == "(hex)" && i > 0 {
			n, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.Itoa(int(n))
			}
			words = remove(words, i)
			i--
		}

		// BIN conversion
		if words[i] == "(bin)" && i > 0 {
			n, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.Itoa(int(n))
			}
			words = remove(words, i)
			i--
		}

		// (up)
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = remove(words, i)
			i--
		}

		// (low)
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = remove(words, i)
			i--
		}

		// (cap)
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = Capitalize(words[i-1])
			words = remove(words, i)
			i--
		}

		// (up, n)
		if strings.HasPrefix(words[i], "(up,") {
			n := extractNumber(words[i])
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = remove(words, i)
			i--
		}

		// (low, n)
		if strings.HasPrefix(words[i], "(low,") {
			n := extractNumber(words[i])
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = remove(words, i)
			i--
		}

		// (cap, n)
		if strings.HasPrefix(words[i], "(cap,") {
			n := extractNumber(words[i])
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = Capitalize(words[i-j])
			}
			words = remove(words, i)
			i--
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(HandleConversions("1E (hex) files were added"))
}
