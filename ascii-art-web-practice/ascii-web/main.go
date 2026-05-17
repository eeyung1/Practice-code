package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler started")

	fmt.Println("Method:", r.Method)

	fmt.Println("Path:", r.URL.Path)

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello World from golang")
}

func main(){
	http.HandleFunc("/", homeHandler)

	fmt.Println("http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}