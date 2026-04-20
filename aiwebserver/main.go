package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// same structs as always
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

// request from browser
type UserMessage struct {
	Message string `json:"message"`
}

// response to browser
type BotResponse struct {
	Response string `json:"response"`
}

// conversation history stored in memory
var (
	history []Message
	mu      sync.Mutex
)

// same chat function as always
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

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AI Coding Mentor</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { 
            font-family: Arial, sans-serif; 
            background: #1a1a2e;
            color: #eee;
            height: 100vh;
            display: flex;
            flex-direction: column;
        }
        header {
            background: #16213e;
            padding: 15px 20px;
            text-align: center;
            font-size: 20px;
            font-weight: bold;
            color: #00d4ff;
            border-bottom: 1px solid #0f3460;
        }
        #chat-box {
            flex: 1;
            overflow-y: auto;
            padding: 20px;
            display: flex;
            flex-direction: column;
            gap: 12px;
        }
        .message {
            max-width: 70%;
            padding: 12px 16px;
            border-radius: 12px;
            line-height: 1.5;
            white-space: pre-wrap;
        }
        .user {
            background: #0f3460;
            align-self: flex-end;
            border-bottom-right-radius: 2px;
        }
        .bot {
            background: #16213e;
            align-self: flex-start;
            border-bottom-left-radius: 2px;
            border: 1px solid #0f3460;
        }
        .thinking {
            color: #888;
            font-style: italic;
        }
        #input-area {
            display: flex;
            padding: 15px;
            background: #16213e;
            border-top: 1px solid #0f3460;
            gap: 10px;
        }
        #user-input {
            flex: 1;
            padding: 12px;
            border-radius: 8px;
            border: 1px solid #0f3460;
            background: #1a1a2e;
            color: #eee;
            font-size: 15px;
            outline: none;
        }
        #user-input:focus {
            border-color: #00d4ff;
        }
        #send-btn {
            padding: 12px 24px;
            background: #00d4ff;
            color: #1a1a2e;
            border: none;
            border-radius: 8px;
            cursor: pointer;
            font-weight: bold;
            font-size: 15px;
        }
        #send-btn:hover {
            background: #00b8d9;
        }
        #send-btn:disabled {
            background: #555;
            cursor: not-allowed;
        }
    </style>
</head>
<body>
    <header>🤖 AI Coding Mentor</header>
    <div id="chat-box">
        <div class="message bot">Hey! I'm your AI coding mentor. Ask me anything about Go, Python, or any programming concept! 🚀</div>
    </div>
    <div id="input-area">
        <input type="text" id="user-input" placeholder="Ask me anything..." autofocus />
        <button id="send-btn" onclick="sendMessage()">Send</button>
    </div>

    <script>
        const chatBox = document.getElementById('chat-box');
        const input = document.getElementById('user-input');
        const btn = document.getElementById('send-btn');

        // send on Enter key
        input.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') sendMessage();
        });

        function addMessage(text, sender) {
            const div = document.createElement('div');
            div.className = 'message ' + sender;
            div.textContent = text;
            chatBox.appendChild(div);
            chatBox.scrollTop = chatBox.scrollHeight;
            return div;
        }

        async function sendMessage() {
            const message = input.value.trim();
            if (!message) return;

            // show user message
            addMessage(message, 'user');
            input.value = '';
            btn.disabled = true;

            // show thinking indicator
            const thinking = addMessage('Thinking...', 'bot thinking');

            try {
                // send to Go server
                const response = await fetch('/chat', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ message: message })
                });

                const data = await response.json();

                // replace thinking with real response
                thinking.textContent = data.response;
                thinking.className = 'message bot';

            } catch (error) {
                thinking.textContent = 'Error connecting to server';
            }

            btn.disabled = false;
            input.focus();
        }
    </script>
</body>
</html>`)
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// read message from browser
	var userMsg UserMessage
	json.NewDecoder(r.Body).Decode(&userMsg)

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not set", 500)
		return
	}

	// add to history with thread safety
	mu.Lock()
	if len(history) == 0 {
		history = append(history, Message{
			Role:    "system",
			Content: "You are a helpful AI coding mentor. Be concise and practical.",
		})
	}
	history = append(history, Message{Role: "user", Content: userMsg.Message})

	// keep history manageable
	if len(history) > 20 {
		history = append(history[:1], history[len(history)-19:]...)
	}
	currentHistory := make([]Message, len(history))
	copy(currentHistory, history)
	mu.Unlock()

	// call Groq
	response, err := chat(apiKey, currentHistory)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// add response to history
	mu.Lock()
	history = append(history, Message{Role: "assistant", Content: response})
	mu.Unlock()

	// send back to browser
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(BotResponse{Response: response})
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/chat", handleChat)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}