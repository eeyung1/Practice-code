package main

import (
	"fmt"
)

func main() {
	lightState := "GREEN" 

	for second := 1; second <= 4; second++ {
		
		fmt.Printf("[Second %d] The light is %s -> ", second, lightState)

		switch lightState {
		case "GREEN":
			fmt.Println("Driver says: 'Go!'")
			lightState = "YELLOW" 

		case "YELLOW":
			fmt.Println("Driver says: 'Slow down!'")
			lightState = "RED"    

		case "RED":
			fmt.Println("Driver says: 'Stop!'")
			lightState = "GREEN"  
		}
	}
}

// func main(){
// 	isLightOn := true

// 	for i := 1; i <= 4; i++ {
// 		isLightOn = !isLightOn

// 		if isLightOn {
// 			fmt.Printf("click %d: The room is bright! (ON)\n", i)
// 		} else {
// 			fmt.Printf("Click %d: The room is dark! (OFF)\n", i)
// 		}
// 	}
// }

// func main() {
// 	inWord := false

// 	for _, ch := range "yes no ok" {
// 		if ch == ' ' {
// 			if !inWord {
// 				fmt.Println("start of word")
// 			}

// 			inWord = true
// 		} else {
// 			inWord = false
// 		}
// 	}
// }

// func main() {
// 	for _, ch := range "a b" {
// 		if ch != ' ' {
// 			fmt.Println("letter")
// 		}
// 	}
// }
