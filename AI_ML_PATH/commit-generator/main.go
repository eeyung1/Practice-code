package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	// Get staged git diff
	cmd := exec.Command("git", "diff", "--staged")
	diffOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running git diff: %v\n", err)
		fmt.Println("Make sure you are in a git repository with staged changes")
		os.Exit(1)
	}

	diff := string(diffOutput)
	if diff == "" {
		fmt.Println("No staged changes found. Stage changes with 'git add' first.")
		os.Exit(1)
	}

	fmt.Printf("📝 Generating commit message for diff (%d bytes)...\n\n", len(diff))

	// Build prompt with few-shot examples
	// Build prompt with few-shot examples - STRICT version
	prompt := fmt.Sprintf(`Generate ONLY the git commit message. No explanations. No commentary. No extra text before or after.

		Format:
		<type>(<scope>): <subject>

		<body>

		Types: feat, fix, docs, style, refactor, test, chore

		Rules:
		- Subject: under 50 characters, imperative mood, no period
		- Body: explain WHAT changed and WHY
		- Use ONLY information from the diff below
		- Do NOT add "For the given diff" or any introductory phrase

		Examples:
		feat(auth): add login with Google OAuth

		- Implement OAuth2 flow for Google
		- Add new endpoint for Google callback

		fix(api): handle null user response

		- Add nil check before dereference
		- Return empty object on error

		Diff:
		%s`, diff)
	// Prepare request
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
		fmt.Printf("API error: %s\n", string(body))
		os.Exit(1)
	}

	var apiResponse map[string]interface{}
	json.Unmarshal(body, &apiResponse)

	content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	fmt.Printf("🤖 Generated commit message:\n\n%s\n\n", content)

	// Ask user for approval
	fmt.Print("Accept? [y/n/edit]: ")
	var answer string
	fmt.Scanln(&answer)

	switch answer {
	case "y", "yes":
		// Commit with generated message
		commitCmd := exec.Command("git", "commit", "-m", content)
		if err := commitCmd.Run(); err != nil {
			fmt.Printf("Error committing: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Committed successfully!")

	case "edit":
		fmt.Print("Edit message: ")
		var edited string
		fmt.Scanln(&edited)
		if edited != "" {
			commitCmd := exec.Command("git", "commit", "-m", edited)
			if err := commitCmd.Run(); err != nil {
				fmt.Printf("Error committing: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Committed with edited message!")
		}

	default:
		fmt.Println("❌ Commit cancelled.")
	}
}
