package main

import (
	"fmt"
	"strconv"
	"strings"
)

func applyHex(text string) string {

	worded := strings.Fields(text)

	for i := 0; i < len(worded); i++ {
		if worded[i] == "(hex)" {
			data, err := strconv.ParseInt(worded[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error")
			}

			worded[i-1] = strconv.FormatInt(data, 10)

			worded = append(worded[:i], worded[i+1:]...)
			i--

		}
	}

	result := strings.Join(worded, " ")
	return result
}

func applyBin(text string) string {

	worded := strings.Fields(text)

	for i := 0; i < len(worded); i++ {
		if worded[i] == "(bin)" {
			data, err := strconv.ParseInt(worded[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error")
			}

			worded[i-1] = strconv.FormatInt(data, 10)

			worded = append(worded[:i], worded[i+1:]...)
			i--

		}
	}

	result := strings.Join(worded, " ")
	return result
}

func applyUp(text string) string {
	worded := strings.Fields(text)

	for i := 0; i < len(worded); i++ {
		if worded[i] == "(up)" {
			worded[i-1] = strings.ToUpper(worded[i-1])
			worded = append(worded[:i], worded[i+1:]...)
			i--
		}
	}
	result := strings.Join(worded, " ")
	return result
}

func applyLow(text string) string {
	worded := strings.Fields(text)

	for i := 0; i < len(worded); i++ {
		if worded[i] == "(low)" {
		worded[i-1] = strings.ToLower(worded[i-1])
		worded = append(worded[:i], worded[i+1:]...)
		i--
		}
	}

	result := strings.Join(worded, " ")
	return result
}

func applyCap(text string) string {
	worded := strings.Fields(text)

	for i := 0; i < len(worded); i++ {
		if worded[i] == "(cap)" {
			worded[i-1] = strings.ToUpper(string(worded[i-1][0])) + worded[i-1][1:]
			worded = append(worded[:i], worded[i+1:]...)
		}
	}

	result := strings.Join(worded, " ")
	return result
}

func editor(text string) string {
	text = applyHex(text)
	text = applyBin(text)
	text = applyUp(text)
	text = applyLow(text)
	text = applyCap(text)
	return text
}