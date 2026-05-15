package main

import (
	"html/template"
	"net/http"
	"ascii-art-web/ascii"
)

type PageData struct {
	AsciiResult string
	Error string
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Template not found", http.StatusInternalServerError)
        return
    }
    
    data := PageData{AsciiResult: "", Error: ""}
    tmpl.Execute(w, data)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    text := r.FormValue("text")
    banner := r.FormValue("banner")
    
    result, genErr := ascii.Generate(text, banner)
    
    tmpl, tmplErr := template.ParseFiles("templates/index.html")
    if tmplErr != nil {
        http.Error(w, "Template not found", http.StatusInternalServerError)
        return
    }
    
    if genErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        data := PageData{AsciiResult: "", Error: genErr.Error()}
        tmpl.Execute(w, data)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    data := PageData{AsciiResult: result, Error: ""}
    tmpl.Execute(w, data)
}