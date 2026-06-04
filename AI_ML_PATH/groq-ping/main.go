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

	requestBody := map[string]interface{}{
		"model":      "llama-3.1-8b-instant",
		"max_tokens": 500,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": "how are tokens calculated in LLM input and output tokens",
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

	fmt.Println("Response:", content)
	fmt.Println("Input tokens:", usage["prompt_tokens"])
	fmt.Println("Output tokens:", usage["completion_tokens"])
}
