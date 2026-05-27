package main

import (
	"fmt"
)

type User struct {
	name    string
	address string
	age     int
	job     string
	salary  int
	status  string
}

func main() {
	var user User

	fmt.Print("Enter name: ")
	fmt.Scanln(&user.name)
	fmt.Print("Enter address: ")
	fmt.Scanln(&user.address)
	fmt.Print("Enter age: ")
	fmt.Scanln(&user.age)
	fmt.Print("Enter job: ")
	fmt.Scanln(&user.job)
	fmt.Print("Enter expected salary in $: ")
	fmt.Scanln(&user.salary)
	fmt.Print("Enter status: Either 'finished' if you are done or 'redo' to correct another information:  ")
	fmt.Scanln(&user.status)

	if user.status == "finished" {
		fmt.Println("Hello", user.name, "You are currently staying at", user.address, "You are", user.age, "years old", "You are currently applying for the position of ", user.job, "and your expected salary is", user.salary)
	}

}

// func main(){
// 	user := User{
// 		address: "Lagos",
// 		age: 35,
// 		job: "AI Engineer",
// 		salary: 30000,
// 	}

// 	fmt.Println("Address is:", user.address)
// 	fmt.Println("Age:", user.age)
// }
