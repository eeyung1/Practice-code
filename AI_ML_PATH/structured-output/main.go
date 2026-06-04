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

	// Read input text
	inputBytes, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Printf("ERROR reading input.txt: %v\n", err)
		os.Exit(1)
	}
	inputText := string(inputBytes)

	// Build prompt asking for JSON
	prompt := fmt.Sprintf(`Extract the following information from this text and return ONLY valid JSON. No other text.

	Text: "%s"

	Return JSON in this exact format:
	{"name": "string", "age": integer, "hobby": "string"}

	If a field is not found, use null.`, inputText)

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
		fmt.Printf("Request error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("API error (status %d): %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var apiResponse map[string]interface{}
	json.Unmarshal(body, &apiResponse)

	// Extract the content string from the response
	content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	fmt.Println("Raw LLM response:")
	fmt.Println(content)
	fmt.Println()

	// Now parse THAT content as JSON
	var person struct {
		Name  string      `json:"name"`
		Age   interface{} `json:"age"` // Use interface{} because age could be number or null
		Hobby string      `json:"hobby"`
	}


	// Try to extract JSON if there's extra text
	cleanJSON := extractJSON(content)
	
	err = json.Unmarshal([]byte(cleanJSON), &person)
	if err != nil {
		fmt.Printf("ERROR parsing JSON from LLM: %v\n", err)
		fmt.Println("Extracted JSON was:", cleanJSON)
		os.Exit(1)
	}

	fmt.Println("Successfully parsed JSON into Go struct:")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %v\n", person.Age)
	fmt.Printf("Hobby: %s\n", person.Hobby)
}

func extractJSON(s string) string {
	start := -1
	end := -1

	for i, c := range s {
		if c == '{' && start == -1 {
			start = i
		}

		if c == '}' {
			end = i
		}
	}

	if start != -1 && end != -1 && end > start {
		return s[start:end+1]
	}

	return s
}
