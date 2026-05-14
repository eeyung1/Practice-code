package main

func MergeBAnners(base map[rune][]string, priority map[rune][]string) map[rune][]string{
	result := make(map[rune][]string)

	for k, v := range base {
		copied := make([]string, len(v))
		copy(copied, v)
		result[k] = copied
	}

	for k, v := range priority {
		copied := make([]string, len(v))
		copy(copied, v)
		result[k] = copied
	}

	return result
}