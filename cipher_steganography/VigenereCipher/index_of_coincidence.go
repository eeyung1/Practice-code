package main

import (
	"unicode"
)

func SplitGroups(ciphertext string, keyLength int) []string {
	groups := make([]string, keyLength)

	for i, ch := range ciphertext {
		groups[i%keyLength] += string(ch)
	}

	return groups
}

func countFrequencies(text string) map[rune]int {
	count := make(map[rune]int)

	for _, ch := range text {
		lower := unicode.ToLower(ch)

		if lower >= 'a' && lower <= 'z' {
			count[lower]++
		}
	}

	return count
}

func IndexOfCoincidence(group string) float64 {
	count := countFrequencies(group)

	n := float64(len(group))

	var ioc float64

	for _, c := range count {
		ioc += float64(c) * (float64(c) - 1)
	}

	ioc = ioc / (n * (n - 1))

	return ioc
}

// func main() {
// 	fmt.Println(SplitGroups("abcdef", 3))
// }
