package main

import (
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")


	if agent == "" {
		agent = "Unknown"
	}

	fmt.Fprintf(w, "You are visiting us using: %s", agent)
}
