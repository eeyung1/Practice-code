package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}

	var result []rune
	for i, r := range str {
		if i % 3 == 2 {
			result = append(result, r)
		}
	}


	return string(result) + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}