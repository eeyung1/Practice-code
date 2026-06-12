package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/download", DownloadHandler)
	fmt.Println("Server started on localhost:8080/hello")
	http.ListenAndServe(":8080", nil)
}



//Quest 5: Receiving Data

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home")

}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="banana.txt`,
	)

	fmt.Fprint(w, text)
}



//Quest 4: Multiple Endpoints
/*
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home")

}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="hello.txt`,
	)

	fmt.Fprint(w, "Hello World")
}
	*/

//Quest 3: Force a Download
// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")

// 	w.Header().Set(
// 		"Content-Disposition",
// 		`attachment; filename="hello.txt`,
// 	)

// 	fmt.Fprint(w, "Hello World")
// }


//Quest 2: Response Headers

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")

// 	fmt.Fprint(w, "Hello World")
// }



//Quest 1 : Create a handler

// func HelloHandler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "Hello World")
// }