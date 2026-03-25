//1. Prints numbers 1 to 10 on one line separated by spaces

/*
package main

import (
	"fmt"
)

func main(){
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	fmt.Println()
}
*/

//2. Prints a 5x5 grid of asterisks using a nested loop

package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}