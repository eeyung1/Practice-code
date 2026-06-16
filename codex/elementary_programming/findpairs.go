package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, target int) string {
	result := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pair := []int{i, j}
				result = append(result, pair)
			}
		}
	}

	if len(result) == 0 {
		return "No pairs found."
	}

	return fmt.Sprintf("Pairs with sum %d: %v", target, result)
}

func main() {
	// Check if exactly 2 arguments are provided (excluding program name)
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	// Validate input format: must start with '[' and end with ']'
	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		fmt.Println("Invalid input.")
		return
	}

	// Remove brackets and split by comma
	trimmed := strings.Trim(arrStr, "[]")
	if trimmed == "" {
		fmt.Println("Invalid input.")
		return
	}

	parts := strings.Split(trimmed, ",")
	arr := []int{}

	// Parse each number
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			fmt.Println("Invalid input.")
			return
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", part)
			return
		}
		arr = append(arr, num)
	}

	// Parse target
	target, err := strconv.Atoi(strings.TrimSpace(targetStr))
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// Find and print pairs
	fmt.Println(findPairs(arr, target))
}