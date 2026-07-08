package main

import "fmt"

func VigenereCipher(text string, keyword string) string {
	shift := []int{}

	for _, ch := range keyword {
		if ch >= 'a' && ch <= 'z' {
			shift = append(shift, int(ch-'a'))
		}
	}

	answer := []rune{}
	keyIndex := 0

	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			value := ch - 'a'
			result := value + rune(shift[keyIndex%len(shift)])
			finalResult := result % 26
			char := finalResult + 'a'
			answer = append(answer, char)
			keyIndex++
		} else if ch >= 'A' && ch <= 'Z' {
			value := ch - 'A'
			result := value + rune(shift[keyIndex%len(shift)])
			finalResult := result % 26
			char := finalResult + 'A'
			answer = append(answer, char)
			keyIndex++
		} else {
			answer = append(answer, ch)
		}
	}

	return string(answer)
}

func main() {
	fmt.Println(VigenereCipher("hello world from goland", "dogz"))
}
