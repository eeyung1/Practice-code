package main

import (
	"fmt"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	if apiKey != "secret123" {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Welcome to the dashboard")
}
