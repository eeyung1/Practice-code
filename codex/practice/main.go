package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func main(){
	colors := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"reset":   "\033[0m",
	}


	activeColor := colors["green"] 
	reset := colors["reset"]
	
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
				result.WriteString(activeColor + bannerLines[index] + reset)
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
