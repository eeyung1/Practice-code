package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main(){
	user := User {
		Name : "John",
		Age : 25,
	}


	ptr := &user

	(*ptr).Name = "Peter"

	fmt.Println(user.Name)

	fmt.Println(ptr)

	fmt.Println(*ptr)

	// fmt.Println(user.Name)

	// ptr.Name = "Peter"

	// fmt.Println(user.Name)

	// fmt.Println(user)
	// fmt.Println(ptr)
	
}