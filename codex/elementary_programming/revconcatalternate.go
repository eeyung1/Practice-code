package main

import (
	"fmt"
)

func RevConcatAlternate(slice1, slice2 []int) []int {
	// Make copies so we don't modify the original slices
	first := make([]int, len(slice1))
	second := make([]int, len(slice2))
	copy(first, slice1)
	copy(second, slice2)

	// Ensure first is the longer slice (or slice1 if equal length)
	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	// Reverse both slices
	for i, j := 0, len(first)-1; i < j; i, j = i+1, j-1 {
		first[i], first[j] = first[j], first[i]
	}

	for i, j := 0, len(second)-1; i < j; i, j = i+1, j-1 {
		second[i], second[j] = second[j], second[i]
	}


	result := []int{}
	lenDiff := len(first) - len(second)


	// Step 1: Add the extra elements from the longer slice (before alternating)
	for i := 0; i < lenDiff; i++ {
		result = append(result, first[i])
	}


	// Step 2: Alternate the remaining elements
	for i := 0; i < len(second); i++ {
		result = append(result, first[lenDiff+i])
		result = append(result, second[i])
	}

	return result
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}