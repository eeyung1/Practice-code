package main

func StackTwo(top []string, bottom []string) []string {
	result := make([]string, len(top)+len(bottom))
	copy(result, top)
	copy(result[len(top):], bottom)
	return result
}

func StackAll(blocks [][]string) []string {
	result := []string{}
	for _, v := range blocks {
		result = StackTwo(result, v)
	}

	return result
}