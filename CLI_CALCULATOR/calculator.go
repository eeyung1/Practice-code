package main

import (
	"fmt"
	
)

func calculate() {
	fmt.Println("welcome to lanta digital calculator, what calculation can we help you with: ")
		var option string
		fmt.Println("addition")
		fmt.Println("subtraction")
		fmt.Println("division")
		fmt.Println("multiplication")
		fmt.Println("modulus")
		fmt.Println("help menu")
	for {
		
		fmt.Println("input an option: ")
		fmt.Scan(&option)
		if option != "addition"&& option != "subtraction"&& option != "division" && option != "multiplication" && option != "modulus" {
			fmt.Println("invalid operation do you need help go to help menu")
			continue
		}

		var digit1 int
		fmt.Println("enter digit1")
		fmt.Scan(&digit1)
		var digit2 int
		fmt.Println("enter digit2")
		fmt.Scan(&digit2)

		if option == "addition" {
			result := digit1 + digit2
			fmt.Printf("%v + %v  =  %v \n", digit1, digit2, result)
		}
		if option == "subtraction" {
			result := digit1 - digit2
			fmt.Printf("%v - %v = %v \n", digit1, digit2, result)
		}
		if option == "division" {
			result := digit1 / digit2
			if digit2 == 0 {
				fmt.Println("cannot divide number by 0")
				continue
			}
			fmt.Printf("%v / %v = %v \n", digit1, digit2, result)
			break
		}
		if option == "multiplication" {
			result := digit1 * digit2
			fmt.Printf("%v * %v = %v \n", digit1, digit2, result)
		}
		if option == "modulus" {
			result := digit1 % digit2
			fmt.Printf("%v modulus %v = %v \n", digit1, digit2, result)
		}
		if option == "help menu" {
			fmt.Println("welcome to lanta calculator this is a digital calculator tht can help you withe the following opration (adition, multiplication, subtraction, division,modulus), of any number all you have to do is to enter any operation of your chioce from the option given input the first and second digit which you want to calculate and it will calculated with the answer for you. thank for using lanta digital calculator!!")
		}
	}
}
func main() {
	calculate()
}
