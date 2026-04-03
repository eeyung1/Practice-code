package main

import (
	"fmt"
)

type Directory struct {
	age   int
	grade string
	city  string
	//school string
}

func main() {
	students := map[string]Directory{
		"Joe":     {age: 15, grade: "S1", city: "Ibadan"},
		"Chaaba":  {age: 17, grade: "S3", city: "Saki"},
		"Mariamu": {age: 12, grade: "J2", city: "Kano"},
		"Chichi":  {age: 16, grade: "S2", city: "Owerri"},
		"Sean":    {age: 17, grade: "S3", city: "Otukpo"},
	}
	students["Musa"] = Directory{age: 14, grade: "S1", city: "Kaduna"}

	students["John"] = Directory{age: 17, grade: "S3", city: "Benue"}
	students["Jonathan"] = Directory{age: 20, grade: "S1", city: "Taraba"}

	chichi := students["Chichi"]
	chichi.grade = "S3"
	students["Chichi"] = chichi

	Joe := students["Joe"]
	Joe.city = "Benue"
	Joe.grade = "S2"
	students["Joe"] = Joe

	Sean :=students["Sean"]
	Sean.age = 11
	Sean.city = "Calabar"
	Sean.grade = "S1"
	students["Sean"] = Sean

	delete(students, "Sean")
	delete(students, "Joe")

	// value, exists := students["John"]
	// if exists {
	// 	fmt.Println("Found:", value)
	// } else {
	// 	fmt.Println("Not Found")
	// }
	for key, value := range students {
		fmt.Println(key, "is", value.age, "years old, in", value.grade, "and lives in", value.city)
		break

		// fmt.Print(students["Joe"])
	}
}
