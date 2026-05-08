package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"ascii-art-web/ascii"
)

// templateData holds all data we pass into the HTML template
type templateData struct {
	Result     string // The generated ASCII art output
	InputText  string // Echo the user's input back into the textarea
	Banner     string // Which banner was selected (to re-select on reload)
	Error      string // Any error message to show the user
}

// loadTemplate parses the index.html template and handles 500 if it fails
func loadTemplate(w http.ResponseWriter) *template.Template {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error: could not load template", http.StatusInternalServerError)
		return nil
	}
	return tmpl
}

// HomeHandler handles GET /
// It renders the main page with empty fields
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET on this route; anything else is a bad request
	if r.Method != http.MethodGet {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	// The root path "/" in Go's default mux matches everything not matched elsewhere.
	// We must explicitly reject paths other than "/"
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	tmpl := loadTemplate(w)
	if tmpl == nil {
		return
	}

	// Render the template with empty/default data
	data := templateData{Banner: "standard"}
	tmpl.Execute(w, data)
}

// AsciiArtHandler handles POST /ascii-art
// It reads form data, generates ASCII art, and renders the result
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	// Parse the incoming form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 - Bad Request: could not parse form", http.StatusBadRequest)
		return
	}

	inputText := r.FormValue("text")
	banner := r.FormValue("banner")

	// --- Input Validation ---

	// Text must not be empty
	if strings.TrimSpace(inputText) == "" {
		tmpl := loadTemplate(w)
		if tmpl == nil {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, templateData{
			Error:  "400 - Bad Request: text input cannot be empty",
			Banner: banner,
		})
		return
	}

	// Banner must be one of the three valid options
	validBanners := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
	if !validBanners[banner] {
		tmpl := loadTemplate(w)
		if tmpl == nil {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, templateData{
			Error:     "400 - Bad Request: invalid banner selected",
			InputText: inputText,
			Banner:    "standard",
		})
		return
	}

	// --- Generate ASCII Art ---
	result, err := ascii.GenerateArt(inputText, banner)
	if err != nil {
		tmpl := loadTemplate(w)
		if tmpl == nil {
			return
		}

		// If banner file wasn't found, that's a 404
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			tmpl.Execute(w, templateData{
				Error:     "404 - Not Found: banner file is missing",
				InputText: inputText,
				Banner:    banner,
			})
			return
		}

		// Otherwise it's a 500
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, templateData{
			Error:     "500 - Internal Server Error: " + err.Error(),
			InputText: inputText,
			Banner:    banner,
		})
		return
	}

	// --- Success: render result on the page ---
	tmpl := loadTemplate(w)
	if tmpl == nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, templateData{
		Result:    result,
		InputText: inputText,
		Banner:    banner,
	})
}
