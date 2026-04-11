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
			num := strings.Trim(words[i+1], ")")
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}

			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}

		// (low, n)
		if strings.HasPrefix(words[i], "(low,") {
			num := strings.Trim(words[i+1], ")")
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}

			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}

		// (cap, n)
		if strings.HasPrefix(words[i], "(cap,") {
			num := strings.Trim(words[i+1], ")")
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}

			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = Capitalize(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func FixPunctuation(text string) string {

	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")

	re2 := regexp.MustCompile(`([.,!?:;])([^\s])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

func FixQuotes(text string) string {

	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}

func FixArticles(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		if strings.ToLower(words[i]) == "a" {

			first := strings.ToLower(string(words[i+1][0]))

			if strings.ContainsAny(first, "aeiouh") {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(HandleConversions("1E (hex) files were added"))
}
