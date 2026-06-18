package main

import (
	"fmt"
	"net/http"
	"strings"
)

func artisttHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")

	fmt.Fprintf(w, "Artist ID: %s", id)
}

func main(){
	http.HandleFunc("/artist/", artisttHandler)

	fmt.Println("Server running at http://localhost:8080/artist/")
	http.ListenAndServe(":8080", nil)
}