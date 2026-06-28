package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/count", countHandler)

	http.HandleFunc("/calculate", CalculateHandler)

	http.HandleFunc("/agent", agentHandler)

	http.HandleFunc("/dashboard", dashboardHandler)

	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/v2", v2Handler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
