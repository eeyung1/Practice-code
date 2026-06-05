package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func UpdateUser(u *User) {
	u.Name = "Peter"
	fmt.Println(*u)
}

func main(){
	user := User{
		Name : "John",
		Age : 25,
	}

	UpdateUser(&user)

	fmt.Println(user.Name)
}