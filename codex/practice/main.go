package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func main(){

	//colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	//colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	colorBlue := "\033[35m"
	colorReset := "\033[0m"
	x := "SCIENCE"

	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bannerLines := strings.Split(string(banner), "\n")

	word := strings.Split(x, "\n")

	var result strings.Builder

	for i := 0; i < 8; i++ {
		for _, line := range word {
			for _, ch := range line {
				index := (int(ch-32)*9) + 1 + i
				result.WriteString(colorBlue + bannerLines[index] + colorReset)
			}
			result.WriteByte('\n')
		}
	}

	fmt.Print(result.String())

}

//PRINT ONLY ONE CHARACTER FROM THE BANNERFILE

// func main() {
// 	banner, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	bannerLines := strings.Split(string(banner), "\n")

// 	for i := 0; i < 8; i++ {
// 		index := ('1'-32)*9 + 1 + i

// 		fmt.Println(bannerLines[index])
// 	}
// }





	// if len(os.Args) != 2 {
	// 	return
	// }


	// for i, r := range os.Args[1] {
	// 	fmt.Println(i, int(r-'0'))
	// }

	// fmt.Println(os.Args[1])
	// fmt.Println([]rune(os.Args[1]))

	// input, _ := strconv.Atoi(os.Args[1])
	// // if err != nil {
	// // 	fmt.Println("Error:", err)
	// // }
	// next, _ := strconv.Atoi(os.Args[2])

	// fmt.Println(input + next)


	// result := input + next
	// ascii := []rune(result)

	// fmt.Println(ascii)
	// fmt.Println(string(ascii))

	//fmt.Println(result)

	//fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1][:1])
