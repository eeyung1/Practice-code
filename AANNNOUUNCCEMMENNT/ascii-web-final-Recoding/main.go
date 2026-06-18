package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	hometmpl, _ := template.ParseFiles("templates/home.html")

	hometmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result := GenerateArt(text, banner)
	data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	tmpl.Execute(w, data)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	if banner == "" {
		banner = r.FormValue("banner")
	}

	result := GenerateArt(text, banner)

	finalResult := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	tmpl.Execute(w, finalResult)
}

func GenerateArt(input string, banner string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	filepath := "banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {
		return "error reading file"
	}

	characters := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(characters, "\n")

	words := strings.Split(input, "\n")

	var result strings.Builder

	for _, word := range words {
		for row := 0; row < 9; row++ {
			for _, char := range word {
				index := int(char-32)*9 + row

				result.WriteString(lines[index])
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/ascii-switch", switchHandler)

	fmt.Println("Server Running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
