package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Cache entry stores a generated message and its diff
type CacheEntry struct {
	Diff    string
	Message string
	Tokens  int
}

// Simple in-memory cache with max 5 entries
type MessageCache struct {
	entries []CacheEntry
	mu      sync.Mutex
}

func (c *MessageCache) Add(diff, message string, tokens int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry := CacheEntry{Diff: diff, Message: message, Tokens: tokens}
	c.entries = append([]CacheEntry{entry}, c.entries...)
	
	// Keep only last 5
	if len(c.entries) > 5 {
		c.entries = c.entries[:5]
	}
}

func (c *MessageCache) Get(diff string) (string, int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	for _, entry := range c.entries {
		if entry.Diff == diff {
			return entry.Message, entry.Tokens, true
		}
	}
	return "", 0, false
}

var cache = &MessageCache{}

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

	// Generate endpoint
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Diff string `json:"diff"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Check cache first
		if msg, tokens, ok := cache.Get(req.Diff); ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":   msg,
				"fromCache": true,
				"tokens":    tokens,
			})
			return
		}

		// Generate new message
		msg, tokens, err := generateCommitMessage(apiKey, req.Diff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Store in cache
		cache.Add(req.Diff, msg, tokens)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   msg,
			"fromCache": false,
			"tokens":    tokens,
		})
	})

	// Regenerate endpoint (forcehttp://localhost:8080 new generation, ignore cache)
	http.HandleFunc("/regenerate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Diff string `json:"diff"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Generate new message (ignore cache)
		msg, tokens, err := generateCommitMessage(apiKey, req.Diff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Update cache with new version
		cache.Add(req.Diff, msg, tokens)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   msg,
			"fromCache": false,
			"tokens":    tokens,
		})
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateCommitMessage(apiKey, diff string) (string, int, error) {
		prompt := fmt.Sprintf(`Generate ONLY the git commit message. No explanations. No commentary. Do NOT copy from the examples.

		The diff below is the ONLY source of truth. Your message MUST describe ONLY what is in this diff.

		Format: <type>(<scope>): <subject>

		<body>

		Types: feat, fix, docs, style, refactor, test, chore

		Rules:
		- Subject under 50 characters, imperative mood
		- Body describes ONLY changes in the diff below
		- Do NOT invent features not in the diff
		- Do NOT repeat the example messages

		Diff (use ONLY this):
		%s`, diff)

		requestBody := map[string]interface{}{
		"model":      "llama-3.1-8b-instant",
		"max_tokens": 200,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", 0, fmt.Errorf("API error: %s", string(body))
	}

	var apiResponse map[string]interface{}
	json.Unmarshal(body, &apiResponse)

	content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	// Extract token usage
	usage := apiResponse["usage"].(map[string]interface{})
	tokens := int(usage["total_tokens"].(float64))

	// Clean up any extra whitespace
	content = strings.TrimSpace(content)

	return content, tokens, nil
}