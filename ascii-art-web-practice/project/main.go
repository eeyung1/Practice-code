package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println(r.FormValue("text"))
	fmt.Println(r.FormValue("banner"))

	tmpl.Execute(w, nil)
}

func main(){
	http.HandleFunc("/", homeHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}