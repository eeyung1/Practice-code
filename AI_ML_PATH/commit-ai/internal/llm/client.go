package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	apiKey     string
	model      string
	maxTokens  int
	httpClient *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GenerateRequest struct {
	Messages []Message
}

type GenerateResponse struct {
	Content string
	Tokens  int
	Error   error
}

// New creates a new LLM client
func New(apiKey, model string, maxTokens int) *Client {
	return &Client{
		apiKey:    apiKey,
		model:     model,
		maxTokens: maxTokens,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GenerateCommitMessage creates a commit message from a git diff
func (c *Client) GenerateCommitMessage(diff string) (string, int, error) {
	prompt := fmt.Sprintf(`Generate ONLY the git commit message. No explanations. No commentary.

The diff below is the ONLY source of truth.

Format: <type>(<scope>): <subject>

<body>

Types: feat, fix, docs, style, refactor, test, chore

Rules:
- Subject under 50 characters, imperative mood
- Body describes ONLY changes in the diff
- Do NOT invent features not in the diff

Diff:
%s`, diff)

	requestBody := map[string]interface{}{
		"model":      c.model,
		"max_tokens": c.maxTokens,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", 0, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", 0, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", 0, fmt.Errorf("parse response: %w", err)
	}

	// Extract content
	choices, ok := apiResponse["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", 0, fmt.Errorf("no choices in response")
	}
	
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("invalid choice format")
	}
	
	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("invalid message format")
	}
	
	content, ok := message["content"].(string)
	if !ok {
		return "", 0, fmt.Errorf("invalid content format")
	}

	// Extract token usage
	usage, ok := apiResponse["usage"].(map[string]interface{})
	if !ok {
		return strings.TrimSpace(content), 0, nil
	}
	
	totalTokens, ok := usage["total_tokens"].(float64)
	if !ok {
		return strings.TrimSpace(content), 0, nil
	}

	return strings.TrimSpace(content), int(totalTokens), nil
}
