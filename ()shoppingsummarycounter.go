package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, char := range str {
		if char == ' ' {
			summary[word]++
			word = ""
		} else {
			word += string(char)
		}
	}
	summary[word]++

	return summary
}
