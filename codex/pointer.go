// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var x int = 42

// 	var p *int = &x

// 	fmt.Println(*p)
// 	*p = 100
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	i, j := 42, 2701

// 	p := &i
// 	fmt.Println(*p)
// 	*p = 21
// 	fmt.Println(i)

// 	p = &j
// 	*p = *p / 37
// 	fmt.Println(j)
// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
