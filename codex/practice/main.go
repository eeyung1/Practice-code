package main

import (
	"fmt"
	"os"
	//"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	for i, r := range os.Args[1] {
		fmt.Println(i, int(r-'0'))
	}

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
}
