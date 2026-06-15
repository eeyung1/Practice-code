package main

import (
	"encoding/json"
	"fmt"
)

type Artist struct {
	Name string `json:"name"`
	Year int 	`json:"year"`
	Age int		`json:"age"`
}

func main() {
	jsonData := `{
		"name":"Queen",
		"year":1970,
		"age":29
	}`

	var artist Artist

	json.Unmarshal([]byte(jsonData), &artist)

	fmt.Println(artist.Name)
	fmt.Println(artist.Year)
	fmt.Println(artist.Age)

	fmt.Println(artist)
}

