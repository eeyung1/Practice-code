package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)

	for i, v := range base {
		result[i] = v
	}

	for a, b := range priority {
		result[a] = b
	}

	return result
}