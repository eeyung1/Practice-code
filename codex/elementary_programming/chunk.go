package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	//if chunk size is 0: print newline and stop
	if size == 0 {
		fmt.Println()
		return
	}

	//create an empty result slice of slices
	result := [][]int{}

	//start from index 0: Move through the slice using chunk-size steps
	for i := 0; i < len(slice); i += size {

		//for each step: set start index
		start := i

		//calculate end index
		end := start + size

		//if end exceeds slice length: set end to slice length
		if end > len(slice) {
			end = len(slice)
		}

		//extract sub-slice using; slice[start:end]
		chunk := slice[start:end]

		//Add sub-slice to result
		result = append(result, chunk)
	}

	//Print final result
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}