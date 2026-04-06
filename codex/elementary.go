package main

import (
	"fmt"
)


func main(){
	
}

/*
func main(){
	seen := make(map[rune]bool)

	input := "aaaccddee"
	result := ""
	count := 0
	for _, r := range input {
		if !seen[r] {
			result += string(r)
			count++
			seen[r] = true
		}
	}

	fmt.Println(result)
	fmt.Println(count)
}

*/


//hello := make(map[rune]bool)
// func main(){


// 	scores := map[string]int{"Alice": 0,}
// 	v := scores["Bob"]
// 	seen := make(map[rune]bool)

// 	fmt.Println(seen['x'])
// 	seen['x'] = true
// 	fmt.Println(!seen['x'])

// 	fmt.Println(v)

// }

/*
func main(){
	make([]int, n)
	v := make([]int, 3, 6)	

	b := make([]int, 0, 3)

	b = append(b, 7)
	b = append(b, 6)
	b = append(b, 5)
	b = append(b, 9)
	fmt.Println(b)
	fmt.Println(v)
}
*/



//slice are windows into an array

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[1:4]

// 	b[0] = 99
// 	fmt.Println(a)
// }
