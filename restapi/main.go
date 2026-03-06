package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is my first Go server!")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintln(w, "Hello stranger!")
		} else {
			fmt.Fprintf(w, "Hello %s!", name)
		}
	})

type User struct {
    Name string `json:"name"`
    City string `json:"city"`
    Age  int    `json:"age"`
}

http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user := User{
        Name: "Alice",
        City: "Lagos",
        Age:  25,
    }
    json.NewEncoder(w).Encode(user)
})

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}