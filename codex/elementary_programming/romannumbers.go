package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	value := os.Args[1]

	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	if num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	calculation := []string{}
	temp := num

	for i := 0; i < len(values); i++ {
		for temp >= values[i] {
			temp -= values[i]
			roman += symbols[i]

			if len(symbols[i]) == 2 {
				calculation = append(calculation, "(" + string(symbols[i][1])+"-"+string(symbols[i][0])+")")
			} else {
				calculation = append(calculation, symbols[i])
			}
		}
	}


	calculationstr := strings.Join(calculation, "+")
	fmt.Println(calculationstr)
	fmt.Println(roman)
}