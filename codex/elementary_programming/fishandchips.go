package main

import (
	"fmt"
)



func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}

	if n%2 == 0 {
		return "fish"
	}

	if n%3 == 0 {
		return "chips"
	}

	return "error: non divisible"
}

func FizzBuzzPlus(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	type rule struct {
		divisor int
		word    string
	}

	rules := []rule{
		{3, "fizz"},
		{5, "buzz"},
		{7, "pop"},
	}

	result := ""

	for _, r := range rules {
		if n%r.divisor == 0 {
			result += r.word
		}
	}

	if result == "" {
		return "error: non divisible"
	}

	return result
}

// func FizzBuzzPlus(n int) string {
// 	if n < 0 {
// 		return "error: number is negative"
// 	}
// 	rules := map[int]string{
// 		3: "fizz",
// 		5: "buzz",
// 		7: "pop",
// 	}

// 	result := ""
// 	for key, value := range rules {
// 		if n % key == 0 {
// 			result += value
// 		}
// 	}

// 	if result == "" {
// 		return "error: non divisible"
// 	}

// 	return result
// }

// func FizzBuzzPlus(n int) string {

// 	if n < 0 {
// 		return "error: number is negative"
// 	}

// 	result := ""

// 	if n%3 == 0 {
// 		result += "fizz"
// 	}

// 	if n%5 == 0 {
// 		result += "buzz"
// 	}

// 	if n%7 == 0 {
// 		result += "pop"
// 	}

// 	if result == "" {
// 		return "error: non divisible"
// 	}

// 	return result
// }

func main() {
	fmt.Println(FizzBuzzPlus(15))
	fmt.Println(FizzBuzzPlus(21))
	fmt.Println(FizzBuzzPlus(35))
	fmt.Println(FizzBuzzPlus(105))
	fmt.Println(FizzBuzzPlus(3))
	fmt.Println(FizzBuzzPlus(5))
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(-1))
	fmt.Println(FishAndChips(6))

}
