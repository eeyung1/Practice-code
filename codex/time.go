package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Welcome to the Playground!")
	fmt.Println("The time is: ", time.Now())

	random_num := rand.Intn(10)

	if random_num%2 == 0 {
		fmt.Println(random_num, "Even")
	} else {
		fmt.Println(random_num, "Odd")
	}

	fmt.Printf("Now you have %g problems.\n",   math.Sqrt(9))
	fmt.Println(math.Pi)

}
