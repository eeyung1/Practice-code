package main

import "fmt"

func main() {
	for i := -5; i <= 5; i++ {
		fmt.Printf("%d %% 26 = %d\n", i, i%26)
	}
}
