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

func chat(apiKey string, messages []Message) (string, error) {
	reqBody := ChatRequest{
		Model:    "llama-3.1-8b-instant",
		Messages: messages,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var chatResp ChatResponse
	json.Unmarshal(body, &chatResp)
	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}
	return chatResp.Choices[0].Message.Content, nil
}

func main() {
	// step 1: validate arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(1)
	}

	// step 2: get API key
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// step 3: read the file
	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// step 4: build messages
	messages := []Message{
		{
			Role:    "system",
			Content: "You are an expert Go code reviewer. Review code for correctness, best practices, potential bugs and improvements. Be specific and practical.",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Please review this Go code and provide feedback:\n\n%s", string(code)),
		},
	}

	// step 5: call chat
	fmt.Println("Reviewing your code...")
	fmt.Println("----------------------")
	response, err := chat(apiKey, messages)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// step 6: print feedback
	fmt.Println(response)
}
