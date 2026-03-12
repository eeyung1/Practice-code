package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	// step 1: validate arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"your question here\"")
		os.Exit(1)
	}

	question := os.Args[1]

	// step 2: get API key
	apiKey := os.Getenv("GROQ_API_KEY")

	if apiKey == "" {
		fmt.Println("Error: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// step 3: build request body
	reqBody := ChatRequest{
		Model: "llama-3.1-8b-instant",
		Messages: []Message{
			{Role: "user", Content: question},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error building request:", err)
		os.Exit(1)
	}

	// step 4: make POST request
	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))

	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// step 5: set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// step 6: send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// step 7: read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	// step 8: parse response
	var chatResp ChatResponse
	json.Unmarshal(body, &chatResp)



	

	// step 9: print answer
	if len(chatResp.Choices) == 0 {
		fmt.Println("No response from API")
		os.Exit(1)
	}

	fmt.Println(chatResp.Choices[0].Message.Content)
}