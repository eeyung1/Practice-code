package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// Conversation history starts empty
	history := []Message{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Conversation Memory Tool ===")
	fmt.Println("Type your message. Commands: :quit, :reset, :show")
	fmt.Println()

	for {
		// Get user input
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Handle commands
		if input == ":quit" {
			fmt.Println("Goodbye!")
			break
		}
		if input == ":reset" {
			history = []Message{}
			fmt.Println("History cleared.")
			continue
		}
		if input == ":show" {
			fmt.Println("\n=== Conversation History ===")
			for i, msg := range history {
				fmt.Printf("%d. %s: %s\n", i+1, msg.Role, msg.Content)
			}
			fmt.Println("===========================")
			continue
		}

		// Add user message to history
		history = append(history, Message{Role: "user", Content: input})

		// Build request with full history
		messages := []map[string]interface{}{}
		for _, msg := range history {
			messages = append(messages, map[string]interface{}{
				"role":    msg.Role,
				"content": msg.Content,
			})
		}

		requestBody := map[string]interface{}{
			"model":      "llama-3.1-8b-instant",
			"max_tokens": 300,
			"messages":   messages,
		}

		jsonBody, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Request error: %v\n", err)
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("API error: %s\n", string(body))
			continue
		}

		var apiResponse map[string]interface{}
		json.Unmarshal(body, &apiResponse)

		content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
		usage := apiResponse["usage"].(map[string]interface{})

		// Add assistant response to history
		history = append(history, Message{Role: "assistant", Content: content})

		fmt.Printf("Assistant: %s\n", content)
		fmt.Printf("[Tokens: in=%v out=%v | History size: %d messages]\n\n",
			usage["prompt_tokens"], usage["completion_tokens"], len(history))
	}
}