package main

import (
	"fmt"
)

func main() {
	result := [][]int{}

	chunk := []int{}

	size := 5

	for i := 1; i <= 100; i++ {
		chunk = append(chunk, i)		
	}

	for j := 0; j < len(chunk); j += size {
		start := j
		end := start + 5

		if end > len(chunk) {
			end = len(chunk)
		}

		Nextchunk := chunk[start:end]

		result = append(result, Nextchunk)
	}


	fmt.Println(result)
}
