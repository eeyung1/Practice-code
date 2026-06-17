package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// type Artist struct {
// 	ID int		`json:"id"`
// 	Name string	`json:"name"`
// }

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func main() {
	resp, err := http.Get(
		"https://groupietrackers.herokuapp.com/api/artists",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var artists []Artist

	err = json.Unmarshal(body, &artists)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(artists[1].Name)

	for _, member := range artists[1].Members {
		fmt.Println(member)
	}

	// for _, artist := range artists {
	// 	fmt.Println(artist.Name)
	// }
}
