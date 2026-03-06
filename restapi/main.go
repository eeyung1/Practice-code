package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "This is my first Go server!")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request){
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintln(w, "Hello stranger!")
		} else {
			fmt.Fprintf(w, "Hello %s!", name)
		}
	})

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
