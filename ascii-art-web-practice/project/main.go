package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	
	result, err := generateAscii(text, banner)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)

		return
	}



	tmpl.Execute(w, result)
}

func generateAscii(text, banner string) (string, error) {
	filepath := banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {

		return "", err
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

	return result, nil
}

func main(){
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiHandler)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}