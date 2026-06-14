package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"ascii-art-web-export/ascii"

)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	art, err := ascii.GenerateAscii(text, banner)

	if err != nil {
		http.Error(w, "404 Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="ascii-art.txt"`,
	)

	w.Header().Set("Content-Length",
		strconv.Itoa(len(art)),
	)
	
	
	fmt.Fprint(w, art)
}