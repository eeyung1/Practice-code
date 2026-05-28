package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)

		return
	}

	result := generateAscii(text, banner)

	tmpl.Execute(w, result)
}

func generateAscii(text, banner string) string {
	filepath := banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {

		return "Error reading banner file"
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")


	lines := strings.Split(content, "\n")

	result := ""

	for i := 0; i < 8; i++ {

		for _, ch := range text {

			index := int(ch-32)*9 + i

			result += lines[index]
		}

		result += "\n"
	}

	return result
}

func main(){
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}