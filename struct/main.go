package main

import (
	"fmt"
)

type Address struct {
	Street string
	City string
}

type User struct {
	Name string
	Age int
	Address Address
}

func main(){
	user := User {
		Name : "John",
		Age : 25,
		Address : Address {
			Street: "12 Allen Avenue",
			City : "Lagos",
		},
	}

	user.Address.City = "Abuja"

	fmt.Println(user.Address.Street)
	fmt.Println(user.Address.City)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Address)
	fmt.Println(user)
}