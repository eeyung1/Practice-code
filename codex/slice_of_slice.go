package main

import (
	"fmt"
)

func main(){
	result := [][]int{}
	counter := 1

	for i := 0; i < 3; i++ {
		row := []int{}

		for j := 0; j < 3; j++ {
			row = append(row, counter)
			counter++
		}

		result = append(result, row)
	}

	for _, row := range result {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}