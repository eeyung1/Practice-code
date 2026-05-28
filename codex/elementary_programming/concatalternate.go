package main

import (
	"fmt"
)

func ConcatAlternate(slice1,slice2 []int) []int {

	//use first for slice1 and second for slice2 to simplify the process
	first := slice1
	second := slice2

	//Decide which slice starts first by comparing slice lengths.
	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	//create empty result slice
	result := []int{}

	//start traversing both slices together
	for i := 0; i < len(second); i++ {

		//At each index:

		//append element from first slice
		result = append(result, first[i])

		//append element from second slice
		result = append(result, second[i])
	}

	//if one slice still has remaining elements: append all remaining values
	result = append(result, first[len(second):]...)

	//return final result
	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}