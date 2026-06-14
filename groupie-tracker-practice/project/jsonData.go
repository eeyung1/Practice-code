package main

import (
	"fmt"
)

type Artist struct{
	Name string
	Year int
}

func main(){

	artist := Artist {
		Name: "Queen",
		Year: 1970,
	}

	fmt.Println(artist.Name)
	fmt.Println(artist.Year)

	fmt.Println(artist)
}
