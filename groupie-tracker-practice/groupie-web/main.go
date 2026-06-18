package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ArtistInfo struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	infos := []ArtistInfo {
		{Name: "Queen"},
		{Name: "ABBA"},
		{Name: "Pink Floyd"},

	}


	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, infos)
}

func main(){
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}