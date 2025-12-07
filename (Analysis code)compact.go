package main

import "fmt"

func Unmatch(c []int) int {
	for i := 0; i < len(c); i++ {
		count := 0
		for j := 0; j < len(c); j++ {
			if c[i] == c[j] {
				count++
			}
		}
		if count%2 != 0 {
			return c[i]
		}
	}
	return -1
}

func main() {
	c := []int{1, 2, 3, 4, 1, 2, 3, 4}
	unmatch := Unmatch(c)
	fmt.Println(unmatch)
}
