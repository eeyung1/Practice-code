package main

import (
	"fmt"
)

func WeAreUnique(str1 , str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	map1 := make(map[rune]bool)
	map2 := make(map[rune]bool)

	for _, r := range str1 {
		map1[r] = true
	}

	for _, ch := range str2 {
		map2[ch] = true
	}

	count := 0

	for r := range map1 {
		if !map2[r] {
			count++
		}
	}

	for r := range map2 {
		if !map1[r] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
