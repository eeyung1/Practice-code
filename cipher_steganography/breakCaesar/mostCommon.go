package main

func mostCommonLetter(freq map[rune]int) rune {
	var mostCommon rune
	maxCount := 0
	for key, value := range freq {
		if value > maxCount {
			maxCount = value
			mostCommon = key
		}

	}

	return mostCommon
}