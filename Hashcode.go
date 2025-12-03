package main

import "fmt"

func HashCode(dec string) string {
	result := ""
	for _, r := range dec {
		if h := (int(r) + len(dec)) % 127; h < 33 {
			result += string(h + 33)
		} else {
			result += string(h)
		}
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
