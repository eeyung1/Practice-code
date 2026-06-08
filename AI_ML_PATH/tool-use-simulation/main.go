package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: GROQ_API_KEY not set")
		os.Exit(1)
	}

	history := []Message{
		{Role: "system", Content: `You are an assistant that can use tools.
	When you need to use a tool, output exactly: <tool>TOOL_NAME:INPUT</tool>
	Available tools:
	- weather:CITY_NAME - Returns weather for a city
	- calculate:EXPRESSION - Performs math (example: calculate:5+3)
	After receiving tool results, continue the conversation normally.
	Do not invent tool results. Always use the tool format when you need information.`},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Tool Use Simulation ===")
	fmt.Println("Available tools: weather (city), calculate (math)")
	fmt.Println("The assistant will use <tool> tags to call them.")
	fmt.Println()

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ":quit" {
			fmt.Println("Goodbye!")
			break
		}

		// Add user message
		history = append(history, Message{Role: "user", Content: input})

		// Loop for multiple rounds of tool calls
		for {
			response := callLLM(apiKey, history)
			fmt.Printf("Assistant: %s\n", response)

			// Extract ALL tool calls from the response
			toolCalls := extractAllToolCalls(response)

			if len(toolCalls) == 0 {
				// No tool calls, add response to history and exit loop
				history = append(history, Message{Role: "assistant", Content: response})
				break
			}

			// Process each tool call
			toolResults := []string{}
			for _, toolCall := range toolCalls {
				fmt.Printf("🔧 TOOL CALL DETECTED: %s\n", toolCall)
				toolResult := executeTool(toolCall)
				fmt.Printf("📊 TOOL RESULT: %s\n", toolResult)
				toolResults = append(toolResults, fmt.Sprintf("%s → %s", toolCall, toolResult))
			}

			// Add assistant's response to history
			history = append(history, Message{Role: "assistant", Content: response})

			// Add combined tool results as a single user message
			combinedResults := strings.Join(toolResults, "\n")
			history = append(history, Message{Role: "user", Content: fmt.Sprintf("Tool results:\n%s", combinedResults)})

			fmt.Println("\n--- Continuing with tool results ---")
		}
	}
}

func callLLM(apiKey string, history []Message) string {
	messages := []map[string]interface{}{}
	for _, msg := range history {
		messages = append(messages, map[string]interface{}{
			"role":    msg.Role,
			"content": msg.Content,
		})
	}

	requestBody := map[string]interface{}{
		"model":      "llama-3.1-8b-instant",
		"max_tokens": 300,
		"messages":   messages,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("API error: %v", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Sprintf("API error: %s", string(body))
	}

	var apiResponse map[string]interface{}
	json.Unmarshal(body, &apiResponse)

	return apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
}

func extractAllToolCalls(response string) []string {
	calls := []string{}
	searchStart := 0
	for {
		start := strings.Index(response[searchStart:], "<tool>")
		if start == -1 {
			break
		}
		start += searchStart
		end := strings.Index(response[start:], "</tool>")
		if end == -1 {
			break
		}
		end += start
		call := response[start+6 : end]
		calls = append(calls, call)
		searchStart = end + 7
	}
	return calls
}

func executeTool(toolCall string) string {
	parts := strings.SplitN(toolCall, ":", 2)
	if len(parts) != 2 {
		return "Error: Invalid tool format. Expected TOOL_NAME:INPUT"
	}

	toolName := parts[0]
	input := parts[1]

	switch toolName {
	case "weather":
		return getWeather(input)
	case "calculate":
		return calculate(input)
	default:
		return fmt.Sprintf("Error: Unknown tool '%s'", toolName)
	}
}

func getWeather(city string) string {
	// Simulated weather data
	weatherData := map[string]string{
		"london":   "🌧️ Rainy, 15°C",
		"paris":    "☀️ Sunny, 22°C",
		"new york": "❄️ Snowy, -2°C",
		"tokyo":    "🌸 Cloudy, 18°C",
	}
	cityLower := strings.ToLower(city)
	if weather, ok := weatherData[cityLower]; ok {
		return weather
	}
	return fmt.Sprintf("☀️ Sunny, 24°C (simulated for %s)", city)
}

func calculate(expression string) string {
	// Very simple calculator
	var a, b int
	var op string
	_, err := fmt.Sscanf(expression, "%d%1s%d", &a, &op, &b)
	if err != nil {
		return fmt.Sprintf("Error: Could not parse '%s'", expression)
	}
	switch op {
	case "+":
		return fmt.Sprintf("%d + %d = %d", a, b, a+b)
	case "-":
		return fmt.Sprintf("%d - %d = %d", a, b, a-b)
	case "*":
		return fmt.Sprintf("%d * %d = %d", a, b, a*b)
	case "/":
		if b == 0 {
			return "Error: Division by zero"
		}
		return fmt.Sprintf("%d / %d = %d", a, b, a/b)
	default:
		return fmt.Sprintf("Error: Unknown operator '%s'", op)
	}
}
