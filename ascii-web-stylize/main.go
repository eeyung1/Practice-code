package main

import (
	"log"
	"net/http"
	"project/handlers"
)

func main(){
	http.HandleFunc("/", handlers.HomeHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/ascii-art", handlers.AsciiHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}



