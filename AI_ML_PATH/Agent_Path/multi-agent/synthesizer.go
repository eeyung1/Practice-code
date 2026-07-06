package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Synthesizer uses LLM to convert raw data into natural language
type Synthesizer struct {
	apiKey string
}

// NewSynthesizer creates a new synthesizer
func NewSynthesizer() *Synthesizer {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Println("WARNING: GROQ_API_KEY not set. Synthesis will use fallback.")
	}
	return &Synthesizer{apiKey: apiKey}
}

// Synthesize converts a result into natural language
func (s *Synthesizer) Synthesize(result Result) string {
	if s.apiKey == "" {
		return s.fallbackSynthesize(result)
	}

	prompt := s.buildPrompt(result)
	response, err := s.callLLM(prompt)
	if err != nil {
		return s.fallbackSynthesize(result)
	}
	return response
}

func (s *Synthesizer) buildPrompt(result Result) string {
	switch result.AgentName {
	case "ArtistAgent":
		artist := result.Data.(Artist)
		return fmt.Sprintf(`Generate a short, natural language summary about this artist:

Name: %s
Members: %v
Formed: %d
First Album: %s

Write a 2-3 sentence summary.`, artist.Name, artist.Members, artist.CreationDate, artist.FirstAlbum)

	case "LocationAgent":
		location := result.Data.(Location)
		return fmt.Sprintf(`Generate a natural language description of these concert locations:

Locations: %v

Write a 1-2 sentence description.`, location.Locations)

	case "DatesAgent":
		date := result.Data.(Date)
		return fmt.Sprintf(`Generate a natural language description of these concert dates:

Dates: %v

Write a 1-2 sentence description.`, date.Dates)

	case "RelationAgent":
		relation := result.Data.(Relation)
		return fmt.Sprintf(`Generate a natural language summary of this artist's tour schedule:

Tour Schedule: %v

Write a 2-3 sentence summary.`, relation.DatesLocations)

	default:
		return fmt.Sprintf("Data: %v", result.Data)
	}
}

func (s *Synthesizer) callLLM(prompt string) (string, error) {
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
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&apiResponse)

	content := apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, nil
}

func (s *Synthesizer) fallbackSynthesize(result Result) string {
	switch result.AgentName {
	case "ArtistAgent":
		artist := result.Data.(Artist)
		return fmt.Sprintf("%s (formed %d) — Members: %v", artist.Name, artist.CreationDate, artist.Members)
	case "LocationAgent":
		location := result.Data.(Location)
		return fmt.Sprintf("Concert locations: %v", location.Locations)
	case "DatesAgent":
		date := result.Data.(Date)
		return fmt.Sprintf("Concert dates: %v", date.Dates)
	case "RelationAgent":
		relation := result.Data.(Relation)
		return fmt.Sprintf("Tour schedule: %v", relation.DatesLocations)
	default:
		return fmt.Sprintf("%v", result.Data)
	}
}