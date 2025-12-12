package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	i, j := 0, 0
	len1, len2 := len(slice1), len(slice2)

	if len1 > len2 {
		for i < len1 || j < len2 {
			if i < len1 {
				result = append(result, slice1[i])
				i++
			}
			if j < len2 {
				result = append(result, slice2[j])
				j++
			}
		}
	} else if len1 < len2 {
		for i < len1 || j < len2 {
			if j < len2 {
				result = append(result, slice2[j])
				j++
			}
			if i < len1 {
				result = append(result, slice1[i])
				i++
			}
		}
	} else {
		for i < len1 || j < len2 {
			if i < len1 {
				result = append(result, slice1[i])
				i++
			}
			if j < len2 {
				result = append(result, slice2[j])
				j++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
