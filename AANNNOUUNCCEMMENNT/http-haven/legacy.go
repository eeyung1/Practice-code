package main

import (
	"fmt"
	"net/http"
)

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/v2", http.StatusMovedPermanently)
}

func v2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to version 2")
}
