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

// package main

// import (
// 	"fmt"
// )

// type student struct {
// 	name string
// 	age int
// 	score int
// }

// func main(){
// 	student1 := student{
// 		name: "James",
// 		age: 25,
// 		score: 70,
// 	}

// 	fmt.Println(student1.name)
// 	fmt.Println(student1.age)
// 	fmt.Println(student1.score)
// }

// package main

// import (
// 	"fmt"
// )

// type student struct {
// 	name string
// 	age int
// 	score int
// }

// func main(){
// 	student1 := student{name: "Jennifer", age: 25, score: 45}
// 	fmt.Printf("Hello %v you are %v years old. Your score is %v\n", student1.name, student1.age, student1.score)
// 	//fmt.Println(student1.name)
// }

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	name string
// 	age int
// 	score int
// }

// func printStudent(s Student) {
// 	fmt.Println(s.name, "is", s.age, "years old and scored", s.score)
// }

// func main(){
// 	student1 := Student{name: "Joel", age: 17, score: 60}
// 	student2 := Student{name: "Joe", age: 19, score: 57}

// 	printStudent(student1)
// 	printStudent(student2)	
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	scores := map[string]int{
// 		"Sean": 75, 
// 		"Tolu": 45,
// 		"Emeka": 90,
// 	}

// 	fmt.Println(scores["Sean"])
// 	fmt.Println(scores["Tolu"])
// 	fmt.Println(scores["Emeka"])
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	capitalCity := map[string]string{
// 		"Nigeria": "Abuja",
// 		"France": "Paris",
// 		"Japan": "Tokyo",
// 	}

// 	for key, value := range capitalCity {
// 		fmt.Println(key, "->", value)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	age int
// 	score int
// }

// func main(){
// 	students := map[string]Student{
// 		"Sean": {age: 20, score: 75},
// 		"Tolu": {age: 22, score: 45},
// 		"Emeka": {age: 19, score: 90},
// 	}

// 	fmt.Println(students["Sean"].score)
// }


package main

import (
	"fmt"
)

type Directory struct {
	age int
	grade string
	city string
}

func main() {
    // 1. Create map first
    students := map[string]Directory{
		"John": {age: 15, grade: "S1", city: "Abuja"},
	}

    // 2. Add new student — OUTSIDE loop
    students["Musa"] = Directory{age: 14, grade: "S1", city: "Kaduna"}

    // 3. Update grade — OUTSIDE loop
    // chichi := students["Chichi"]
    // chichi.grade = "S3"
    // students["Chichi"] = chichi

    // 4. Delete — OUTSIDE loop
    // delete(students, "Sean")

    // 5. Check if exists — OUTSIDE loop
    value, exists := students["Sean"]
    if exists {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Not Found!")
    }

    // 6. Loop LAST — just for printing
    for key, value := range students {
        fmt.Println(key, "is", value.age, "years old, in", value.grade, "and lives in", value.city)
    }
}