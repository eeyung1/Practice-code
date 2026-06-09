package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Title  string
	Result string
}

var tmpl = template.Must(template.ParseFiles("template/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	Data := pageData{

		Title:  "Ascii-Art-Web",
		Result: "",
	}

	tmpl.Execute(w, Data)

}
func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Page Not Found", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodPost {

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result := artbuilder(text , banner)

	Data := pageData{
		Title:  "Ascii-Art-Web",
		Result: result,
	}
	tmpl.Execute(w, Data)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("server is live on port 8000...")
	http.ListenAndServe(":8000", nil)
}
