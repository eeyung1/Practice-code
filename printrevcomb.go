package main

import (
	"fmt"
)

func main() {
	firstRun := true

	for d1 := 9; d1 >= 2; d1-- {
		for d2 := d1 - 1; d2 >= 1; d2-- {
			for d3 := d2 - 1; d3 >= 0; d3-- {
				if !firstRun {
					fmt.Print(", ")
				}

				fmt.Printf("%d%d%d", d1, d2, d3)
				firstRun = false
			}
		}
	}
	fmt.Println()
}
