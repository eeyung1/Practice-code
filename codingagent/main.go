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

func extractCode(response string) string {
	// extract code from markdown code blocks
	if strings.Contains(response, "```go") {
		parts := strings.Split(response, "```go")
		if len(parts) > 1 {
			code := strings.Split(parts[1], "```")[0]
			return strings.TrimSpace(code)
		}
	}
	if strings.Contains(response, "```") {
		parts := strings.Split(response, "```")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}
	return strings.TrimSpace(response)
}

func runAgent(apiKey string, task string) (string, error) {
	history := []Message{
		{
			Role: "system",
			Content: `You are an expert Go programmer. 
When given a task, write complete working Go code.
When given an error, fix the code.
IMPORTANT RULES:
- Always return ONLY the complete Go code
- Always include package main and imports
- Always include a main() function that demonstrates the solution
- Never include explanations outside the code
- Use comments inside the code to explain`,
		},
	}

	currentTask := task
	maxAttempts := 5

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("\n🔄 Attempt %d/%d\n", attempt, maxAttempts)
		fmt.Println("📝 Asking AI to write code...")

		// add task to history
		history = append(history, Message{
			Role:    "user",
			Content: currentTask,
		})

		// get code from AI
		response, err := chat(apiKey, history)
		if err != nil {
			return "", err
		}

		// add response to history
		history = append(history, Message{
			Role:    "assistant",
			Content: response,
		})

		// extract clean code
		code := extractCode(response)

		fmt.Println("💾 Saving code to temp file...")
		os.WriteFile("temp_agent.go", []byte(code), 0644)

		fmt.Println("🚀 Running code...")
		output, err := exec.Command("go", "run", "temp_agent.go").CombinedOutput()

		if err == nil {
			// success!
			fmt.Println("\n✅ Code works!")
			fmt.Println("\n📤 Output:")
			fmt.Println(string(output))
			fmt.Println("\n💻 Final Code:")
			fmt.Println(code)
			os.Remove("temp_agent.go")
			return code, nil
		}

		// failed — show error and try again
		errorMsg := string(output)
		fmt.Printf("\n❌ Error on attempt %d:\n%s\n", attempt, errorMsg)

		if attempt < maxAttempts {
			fmt.Println("🔧 Asking AI to fix the error...")
			currentTask = fmt.Sprintf(`The code you wrote has an error. Fix it.

ERROR:
%s

YOUR PREVIOUS CODE:
%s

Return the complete fixed Go code only.`, errorMsg, code)
		}
	}

	os.Remove("temp_agent.go")
	return "", fmt.Errorf("failed after %d attempts", maxAttempts)
}

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: GROQ_API_KEY not set")
		os.Exit(1)
	}

	fmt.Println("🤖 Self-Improving Coding Agent")
	fmt.Println("================================")

	// test tasks
	tasks := []string{
		"Write a Go function that reverses a string and print the result",
		"Write a Go function that checks if a number is prime and test it with numbers 1-20",
		"Write a Go function that calculates fibonacci sequence up to 10 numbers",
	}

	for _, task := range tasks {
		fmt.Printf("\n📋 TASK: %s\n", task)
		fmt.Println(strings.Repeat("-", 50))

		_, err := runAgent(apiKey, task)
		if err != nil {
			fmt.Println("❌ Agent failed:", err)
		}

		fmt.Println(strings.Repeat("=", 50))
	}
}
