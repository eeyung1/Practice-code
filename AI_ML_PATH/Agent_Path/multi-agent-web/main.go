package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "strconv"
)

func main() {
	// Serve static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	// API endpoint
	http.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			ArtistID int    `json:"artistID"`
			Query    string `json:"query"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Select agents based on query
		orchestrator := NewOrchestrator()
		agents := orchestrator.SelectAgents(req.Query, req.ArtistID)

		// Run agents
		coordinator := Newcoordinator(agents)
		results := coordinator.Run()

		// Synthesize results
		synthesizer := NewSynthesizer()
		var responseResults []map[string]interface{}

		for _, result := range results {
			resp := map[string]interface{}{
				"agent": result.AgentName,
			}

			if result.Error != nil {
				resp["error"] = result.Error.Error()
			} else {
				resp["response"] = synthesizer.Synthesize(result)
			}

			responseResults = append(responseResults, resp)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"results": responseResults,
		})
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
