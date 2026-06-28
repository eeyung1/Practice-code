package main

import (
	"fmt"
	"net/http"
)



func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	text := r.URL.Query().Get("name")

	if text == "" {
		text = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", text)
}