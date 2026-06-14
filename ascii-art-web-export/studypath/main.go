//[/download?text=HELLO&banner=shadow]

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/download", DownloadHandler)
	fmt.Println("Server started on localhost:8080/hello")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home")

}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	ascii, err := GenerateAscii(text, banner)

	if err != nil {
		http.Error(w, "404 Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="ascii-art.txt`,
	)

	w.Header().Set("Content-Length",
		strconv.Itoa(len(ascii)),
	)
	
	fmt.Fprint(w, ascii)
}

func GenerateAscii(text, banner string) (string, error) {
	// good := []rune(text)

	// fmt.Println(good)

	filepath := "../banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {

		return "", err
	}

	if strings.Contains(text, `\\n`) {
		text = strings.ReplaceAll(text, `\\n`, "\n")
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(content, "\n")

	characters := strings.ReplaceAll(text, "\r\n", "\n")

	if strings.Contains(text, `\n`) {
		characters = strings.ReplaceAll(text, `\n`, "\n")
	}

	parts := strings.Split(characters, "\n")

	result := ""

	for _, part := range parts {

		for i := 0; i < 8; i++ {

			for _, ch := range part {

				index := int(ch-32)*9 + i

				result += lines[index]
			}

			result += "\n"
		}

		result += "\n"
	}

	return result, nil
}

// Quest 6: Receive Multiple Inputs

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome Home")

// }
// func DownloadHandler(w http.ResponseWriter, r *http.Request) {
// 	text := r.URL.Query().Get("text")
// 	banner := r.URL.Query().Get("banner")

// 	w.Header().Set("Content-Type", "text/plain")

// 	w.Header().Set(
// 		"Content-Disposition",
// 		`attachment; filename="info.txt`,
// 	)

// 	fmt.Fprintf(
// 		w,
// 		"Text: %s\nBanner: %s",
// 		text,
// 		banner,
// 	)
// }

// //Quest 5: Receiving Data

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome Home")

// }
// func DownloadHandler(w http.ResponseWriter, r *http.Request) {
// 	text := r.URL.Query().Get("text")
// 	w.Header().Set("Content-Type", "text/plain")

// 	w.Header().Set(
// 		"Content-Disposition",
// 		`attachment; filename="banana.txt`,
// 	)

// 	fmt.Fprint(w, text)
// }

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
