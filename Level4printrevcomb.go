package main

import "fmt"

func main() {
	combinations := []string{}

	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				combination := fmt.Sprintf("%d%d%d", i, j, k)
				combinations = append(combinations, combination)
			}
		}
	}

	result := fmt.Sprintf("%s", combinations[0])
	for _, comb := range combinations[1:] {
		result += ", " + comb
	}

	fmt.Println(result)
}
