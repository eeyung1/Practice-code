package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"project/ascii"
)


func AsciiHandler(w http.ResponseWriter, r *http.Request) {
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
	
	result, err := ascii.GenerateAscii(text, banner)
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



	if err = tmpl.Execute(w, result); err != nil {
		log.Println("template execute error:", err)
	}

}