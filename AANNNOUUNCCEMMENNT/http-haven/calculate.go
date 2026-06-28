package main

import (
	"fmt"
	"strconv"
	"net/http"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	opp := r.URL.Query().Get("op")
	num1 := r.URL.Query().Get("a")
	num2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(num1)

	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(num2)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	switch opp {
	case "add":
		fmt.Fprintf(w, "Result: %d", a+b)
		// fallthrough
	case "subtract":
		fmt.Fprintf(w, "Result: %d", a-b)
	case "multiply":
		fmt.Fprintf(w, "Result: %d", a*b)
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	}

}