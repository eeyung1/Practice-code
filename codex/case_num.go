/*Question 3 - Multi-Word Commands
Task:
Support commands like:
(up, 2)
(low, 3)
(cap, 4)
Example:
["this", "is", "so", "fun", "(up,", "2)"]
Output:
["this", "is", "SO", "FUN"]
Requirements:
Detect (command, n)
Apply to last n words
Handle edge cases:
n > number of words
invalid numbers
 Bonus:
 Combine with single commands:
(up) vs (up, 2)
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func conv(words []string) []string {
	car := ""
	word := ""
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up,") && i != 0 {
			for _, r := range words[i] {
				if r >= '0' && r <= '9' {
					car += string(r)
				}
			}
			a, _ := strconv.Atoi(car)
			j := i - 1
			if a > j {
				a = j
			}
			for j = i - a - 1; j <= i-1; j++ {
				words[j] = strings.ToUpper(words[j])
			}
		}
		if strings.Contains(words[i], "(low,") && i != 0 {
			for _, r := range words[i] {
				if r >= '0' && r <= '9' {
					car += string(r)
				}
			}
			a, _ := strconv.Atoi(car)
			j := i - 1
			if a > j {
				a = j
			}
			for j = i - a - 1; j <= i-1; j++ {
				words[j] = strings.ToLower(words[j])
			}
		}
		if strings.Contains(words[i], "(cap,") && i != 0 {
			for _, r := range words[i] {
				if r >= '0' && r <= '9' {
					car += string(r)
				}
			}
			a, _ := strconv.Atoi(car)
			j := i - 1
			if a > j {
				a = j
			}
			for j = i - a - 1; j <= i-1; j++ {

				words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(string(words[j][1:]))
			}
		}
	}
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up,") || strings.Contains(words[i], "(low,") || strings.Contains(words[i], "cap,") {
			word += ""
		} else {
			word += words[i] + " "
		}
	}
	return strings.Fields(word)
}

func main() {
	fmt.Println(conv([]string{"(up, 4)", "this", "is", "so", "fun", "you"}))
	fmt.Println(conv([]string{"this", "is", "SO", "FUN", "(low, 2)", "you"}))
	fmt.Println(conv([]string{"this", "is", "SO", "FUN", "(cap, 2)", "you"}))
}
