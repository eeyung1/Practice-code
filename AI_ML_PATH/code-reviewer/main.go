package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// Check command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.go>")
		os.Exit(1)
	}
	filename := os.Args[1]

	// Read the Go file
	codeBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	code := string(codeBytes)

	// Build the prompt
	prompt := fmt.Sprintf(`You are a senior Go code reviewer. Review the following Go code for:

	1. Unused variables or imports
	2. Missing error handling
	3. Potential bugs or edge cases
	4. Code style issues (naming, formatting)
	5. Performance concerns

	Return your review in this exact format:

	CRITICAL ISSUES:
	- (list critical problems that would cause bugs)

	WARNINGS:
	- (list potential issues or bad practices)

	SUGGESTIONS:
	- (list improvements for readability or performance)

	Code to review:

	%s`, code)

	// Prepare request
	requestBody := map[string]interface{}{
		"model":      "llama-3.1-8b-instant",
		"max_tokens": 800,
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
		fmt.Printf("Request error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("API error: %s\n", string(body))
		os.Exit(1)
	}

	var apiResponse map[string]interface{}
	json.Unmarshal(body, &apiResponse)

	content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	fmt.Println("=== CODE REVIEW REPORT ===")
	fmt.Println(content)
	fmt.Println("\n=== END OF REPORT ===")
}