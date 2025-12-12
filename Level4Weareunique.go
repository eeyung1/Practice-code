package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	for _, char := range str1 {
		set1[char] = true
	}

	for _, char := range str2 {
		set2[char] = true
	}

	uniqueCount := 0

	for char := range set1 {
		if !set2[char] {
			uniqueCount++
		}
	}

	for char := range set2 {
		if !set1[char] {
			uniqueCount++
		}
	}

	return uniqueCount
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
