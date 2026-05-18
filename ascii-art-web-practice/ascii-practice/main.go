package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	html := `
	<form action="/submit" method="POST">

		<textarea name="text"></textarea>

		<button type="submit">Send</button>

	</form>
	`

	fmt.Fprint(w, html)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	fmt.Println("Text:", text)
	fmt.Fprint(w, text)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	fmt.Println("server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
