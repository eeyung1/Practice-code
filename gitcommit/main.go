package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

func getDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(string(output)) == "" {
		cmd = exec.Command("git", "diff")
		output, err = cmd.Output()
		if err != nil {
			return "", err
		}
	}
	if strings.TrimSpace(string(output)) == "" {
		return "", fmt.Errorf("no changes found")
	}
	return string(output), nil
}

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: GROQ_API_KEY not set")
		os.Exit(1)
	}

	diff, err := getDiff()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	messages := []Message{
		{
			Role:    "system",
			Content: "You are a git commit message generator. Generate concise, descriptive commit messages following conventional commits format: type(scope): description. Types: feat, fix, docs, refactor, test. One line only, max 72 chars.",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Generate a commit message for these changes:\n\n%s", diff),
		},
	}

	fmt.Println("Analyzing changes...")
	response, err := chat(apiKey, messages)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("\nSuggested commit message:")
	fmt.Println("→", response)

	fmt.Print("\nApprove? (y/n/e to edit): ")
	var input string
	fmt.Scanln(&input)

	switch strings.ToLower(input) {
	case "y":
		cmd := exec.Command("git", "commit", "-m", response)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		fmt.Println("✅ Committed!")
	case "e":
		fmt.Print("Enter your message: ")
		var custom string
		fmt.Scanln(&custom)
		cmd := exec.Command("git", "commit", "-m", custom)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		fmt.Println("✅ Committed!")
	default:
		fmt.Println("Commit cancelled")
	}
}