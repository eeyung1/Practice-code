package main

import (
	"ascii-art-web-export/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}
