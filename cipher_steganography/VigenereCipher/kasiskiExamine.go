package main

import "fmt"

func kasiskiExamine(cipherText string) map[string][]int {
	positions := make(map[string][]int)

	for seqLen := 3; seqLen <= 5; seqLen++ {
		for i := 0; i < len(cipherText)-seqLen; i++ {
			sequence := cipherText[i : i+seqLen]
			positions[sequence] = append(positions[sequence], i)
		}
	}

	for sequence, pos := range positions {
		if len(pos) < 2 {
			delete(positions, sequence)
		}
	}

	return positions
}

func main() {
	plaintext := "the quick brown fox jumps over the lazy dog the quick brown fox jumps over the lazy dog"
	cipherText := VigenereCipher(plaintext, "dogz")

	fmt.Println(cipherText)

	fmt.Println(kasiskiExamine(cipherText))

}
