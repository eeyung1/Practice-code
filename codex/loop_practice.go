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

/*
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

*/

//3. Counts DOWN from 10 to 1 then prints "Blastoff!"
/*

package main

import (
	"fmt"
)

func main(){
	for i := 10; i >= 1; i-- {
		
		if i == 1 {
			fmt.Println("Blastoff!")
		} else {
			fmt.Println(i)
		}

	}
}

*/
//4. Iterates over the string "GoLang" and prints each character
//   on its own line with its index and ASCII value

/*

package main

import (
	"fmt"
)

func main(){
	word := "Golang"

	for i, r := range word {
		fmt.Printf("%d %c = %d\n", i, r, r)
	}
}
	
*/

//Keeps doubling a number starting from 1, stopping
  // when it exceeds 1000, prints each value


package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 1000; i *= 2 {
		fmt.Println(i)
	}
}