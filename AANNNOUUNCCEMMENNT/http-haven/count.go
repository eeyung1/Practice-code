package main

import (
	"fmt"
	"io"
	"net/http"
)


func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request")
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer r.Body.Close()

	fmt.Fprint(w, len(data))
}