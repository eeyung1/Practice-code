package main

import (
	"fmt"
)

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

func getDistances(positions map[string][]int) []int {
	result := []int{}

	for _, position := range positions {
		for i := 1; i < len(position); i++ {
			result = append(result, position[i]-position[i-1])
		}
	}

	return result
}

func getFactors(num int) []int {
	factors := []int{}

	for i := 2; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func mostCommonFactor(distance []int) map[int]int {

	freqCount := map[int]int{}

	for _, num := range distance {
		factors := getFactors(num)

		for _, factor := range factors {
			freqCount[factor]++
		}
	}

	return freqCount
}

func likelyKeyLength(mostCommon map[int]int) int {
	bestFactor := 0
	bestCount := 0
	for factor, count := range mostCommon {
		if factor <= 13 {
			if count > bestCount {
				bestCount = count
				bestFactor = factor
			}
		}
	}

	return bestFactor
}

func main() {
	plaintext := "the quick brown fox jumps over the lazy dog the quick brown fox jumps over the lazy dog"
	cipherText := VigenereCipher(plaintext, "dogz")

	fmt.Println(cipherText)

	positions := kasiskiExamine(cipherText)
	distance := getDistances(positions)
	factors := mostCommonFactor(distance)


	fmt.Println(positions)
	fmt.Println(distance)
	fmt.Println(factors)
	fmt.Println(likelyKeyLength(factors))

}
