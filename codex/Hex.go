package main

import (
	"fmt"
	"strings"
	"strconv"
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

func applBin(text string) string {

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



// func main(){
// 	fmt.Println(applyHex("1E (hex) files were added"))
// 	fmt.Println(applBin("It has been 10 (bin) years"))
// }
