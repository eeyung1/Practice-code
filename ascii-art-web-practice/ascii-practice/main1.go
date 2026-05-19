package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Println("PATH:", r.URL.Path)
	fmt.Fprint(w, "HOME PAGE")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ABOUT PAGE")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is used for testing the page")
}

func main(){
	http.HandleFunc("/", homeHandler)
	
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server starded on http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}