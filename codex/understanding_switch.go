// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var day string
// 	fmt.Print("Input a day: ")
// 	fmt.Scanln(&day)
// 	switch day {
// 	case "Monday":
// 		fmt.Println("Start of the week!")
// 	case "Friday":
// 		fmt.Println("Almost the weekend!")
// 	case "Saturday", "Sunday":
// 		fmt.Println("It's the weekend!")
// 	default:
// 		fmt.Println("It's a regular weekday")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var score int
// 	fmt.Print("Input score: ")
// 	fmt.Scan(&score)

// 	switch {
// 	case score >= 70:
// 		fmt.Println("A")
// 	case score >= 60:
// 		fmt.Println("B")
// 	case score >= 50:
// 		fmt.Println("C")
// 	default:
// 		fmt.Println("F - Failed")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var greetings string

// 	fmt.Print("Input the time of the day either morning/afternoon/evening: ")
// 	fmt.Scan(&greetings)

// 	switch {
// 	case greetings == "morning":
// 		fmt.Println("Good Morning!")
// 	case greetings == "afternoon":
// 		fmt.Println("Good Afrernoon!")
// 	case greetings == "evening":
// 		fmt.Println("Godd Evening!")
// 	default:
// 		fmt.Println("Hello!")
// 	}
// }

package main

import (
	"fmt"
)

type student struct {
	name string
	age int
	score int
}

func main(){
	student1 := student{
		name: "James",
		age: 25,
		score: 70,
	}

	fmt.Println(student1.name)
	fmt.Println(student1.age)
	fmt.Println(student1.score)
}