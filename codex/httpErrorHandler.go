package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ok" {
		http.Error(w, "200 Ok", http.StatusOK)
		return
	}

	if r.URL.Path == "/notfound" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.URL.Path == "/badrequest" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/error" {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/" {
		http.Error(w, "404 Not found Available path /ok, /badrequest, /notfound, /error", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 Not found Available path /ok, /badrequest, /notfound, /error", http.StatusNotFound)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website")
}

// func OkHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "200 OK", http.StatusOK)
// 		return
// 	}
// }

// func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "400 Bad Request", http.StatusBadRequest)
// 		return
// 	}
// }

// func ErrorHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "500 Internal Server Error", http.StatusBadRequest)
// 		return
// 	}
// }

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	// http.HandleFunc("/ok", OkHandler)
	// http.HandleFunc("/badrequest", BadRequestHandler)
	// http.HandleFunc("/error", ErrorHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
