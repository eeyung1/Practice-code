package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	salary int
}

func main() {
	var pers1 person

	pers1.name = "John"
	pers1.age = 34
	pers1.salary = 45000

	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("salary: ", pers1.salary)
}
