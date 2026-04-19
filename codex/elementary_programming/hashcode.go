package main

import (
	"fmt"
)

func HashCode(dec string) string {

	//Algorithm
	// create an empty result
	// loop through characters
	// hash each character
	// fix unprintable chars
	// append to result
	// return final string

	if dec == "" {
		return ""
	}

	var result []rune
	size := rune(len(dec))

	//fmt.Println(size)

	for _, r := range dec {
		hashed := (r + size) % 127

		if hashed < 32 {
			hashed += 33
		}

		result = append(result, hashed)
	}

	return string(result)
}

func DecodeHash(enc string) string {
	if enc == "" {
		return ""
	}

	var result []rune
	size := rune(len(enc))

	for _, r := range enc {
		Dehashed := (r - size + 127) % 127
		result = append(result, Dehashed)
	}

	return string(result)
}

func ShiftCipher(s string, shift int) string{
	if s == "" {
		return ""
	}

	var result []rune
	size := rune(shift)

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			shifted := ((r - 'a') + size) % 26 + 'a'
			result = append(result, shifted)
		} else {
			result = append(result, r)
		}
	}


	return string(result)
}

func main() {
	fmt.Println(ShiftCipher("abc", 2))
	fmt.Println(ShiftCipher("xyz", 2))
	fmt.Println(DecodeHash("CD"))
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
