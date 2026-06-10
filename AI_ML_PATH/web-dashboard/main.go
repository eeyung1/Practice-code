package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// Serve static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle chat with streaming
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Prompt string `json:"prompt"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Set up streaming response
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		// Prepare request to Groq
		requestBody := map[string]interface{}{
			"model":      "llama-3.1-8b-instant",
			"max_tokens": 500,
			"stream":     true, // Enable streaming
			"messages": []map[string]interface{}{
				{
					"role":    "user",
					"content": req.Prompt,
				},
			},
		}

		jsonBody, _ := json.Marshal(requestBody)

		httpReq, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		resp, err := client.Do(httpReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

				// Read streaming response
		scanner := bufio.NewScanner(resp.Body)
		var fullResponse strings.Builder

		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data: ") {
				continue
			}
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}

			var chunk map[string]interface{}
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				continue
			}

			// Extract content from chunk
			choices, ok := chunk["choices"].([]interface{})
			if !ok || len(choices) == 0 {
				continue
			}
			delta, ok := choices[0].(map[string]interface{})["delta"].(map[string]interface{})
			if !ok {
				continue
			}
			content, ok := delta["content"].(string)
			if !ok || content == "" {
				continue
			}

			// Write to client
			fmt.Fprint(w, content)
			fullResponse.WriteString(content)
			flusher.Flush()
		}

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(w, "\n\n[Error reading stream: %v]", err)
			flusher.Flush()
		}
		// Note: Token counts not available in stream mode from this simple implementation
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}