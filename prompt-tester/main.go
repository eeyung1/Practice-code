package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	promptFile := "prompt.txt"
	if len(os.Args) > 1 {
		promptFile = os.Args[1]
	}

	// Read prompt from file
	promptBytes, err := os.ReadFile(promptFile)
	if err != nil {
		fmt.Printf("ERROR reading %s: %v\n", promptFile, err)
		os.Exit(1)
	}
	prompt := string(promptBytes)

	requestBody := map[string]interface{}{
		"model":      "llama-3.1-8b-instant",
		"max_tokens": 500,
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
		fmt.Printf("API error (status %d): %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	content := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	usage := result["usage"].(map[string]interface{})

	// Save response to file with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("response_%s.txt", timestamp)
	saveData := fmt.Sprintf("PROMPT:\n%s\n\n---\nRESPONSE:\n%s\n\n---\nTOKENS:\nInput: %v\nOutput: %v\n",
		prompt, content, usage["prompt_tokens"], usage["completion_tokens"])

	err = os.WriteFile(filename, []byte(saveData), 0644)
	if err != nil {
		fmt.Printf("Warning: Could not save response: %v\n", err)
	}
	fmt.Printf("Saved to: %s\n", filename)

	fmt.Println("Response:", content)
	fmt.Println("Input tokens:", usage["prompt_tokens"])
	fmt.Println("Output tokens:", usage["completion_tokens"])
}
