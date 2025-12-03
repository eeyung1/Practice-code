package main

import "fmt"

func RepeatAlpha(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count := int(r - 'a' + 1)
			for i := 0; i < count; i++ {
				res += string(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			count := int(r - 'A' + 1)
			for i := 0; i < count; i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
