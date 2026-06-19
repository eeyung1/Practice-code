package main

import (
	"fmt"
	"io"
	"strconv"

	// "io"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	text := r.URL.Query().Get("name")

	if text == "" {
		text = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", text)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request")
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer r.Body.Close()

	fmt.Fprint(w, len(data))
}

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
		fallthrough
	case "subtract":
		fmt.Fprintf(w, "Result: %d", a-b)
	case "multiply":
		fmt.Fprintf(w, "Result: %d", a*b)
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	}

}

func agentHandler(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")

	fmt.Println(agent)

	if agent == "" {
		agent = "Unknown"
	}

	fmt.Fprintf(w, "You are visiting us using: %s", agent)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	if apiKey != "secret123" {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Welcome to the dashboard")
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/count", countHandler)

	http.HandleFunc("/calculate", CalculateHandler)

	http.HandleFunc("/agent", agentHandler)

	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
