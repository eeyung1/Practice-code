package main

import (
	"fmt"
)

func ConcatSlice(slice1, slice2 []int) []int {
	result := []int{}

	for _, elem := range slice1 {
		result = append(result, elem)
	}

	for _, elem := range slice2 {
		result = append(result, elem)
	}

	return result
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}