package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	
	"commit-ai/internal/cache"
	"commit-ai/internal/config"
	"commit-ai/internal/history"
	"commit-ai/internal/llm"
)

var (
	configPath  = flag.String("config", "config.json", "path to config file")
	webDir      = flag.String("web", "web", "path to web directory")
)

func main() {
	flag.Parse()
	
	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	// Initialize components
	llmClient := llm.New(cfg.APIKey, cfg.Model, cfg.MaxTokens)
	memCache := cache.New(cfg.MaxCacheSize, 60) // 60 minute TTL
	
	// Initialize history database
	dbPath := filepath.Join("data", "history.db")
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Printf("Warning: could not create data directory: %v", err)
	}
	
	db, err := history.New(dbPath)
	if err != nil {
		log.Printf("Warning: could not open history database: %v", err)
		log.Println("Continuing without history persistence")
	}
	if db != nil {
		defer db.Close()
	}
	
	// Serve static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(*webDir, "index.html"))
	})
	
	// Generate endpoint
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		var req struct {
			Diff string `json:"diff"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		
		// Check memory cache
		if msg, tokens, ok := memCache.Get(req.Diff); ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":   msg,
				"fromCache": true,
				"tokens":    tokens,
			})
			
			// Save to history
			if db != nil {
				db.Save(req.Diff, msg, tokens, true)
			}
			return
		}
		
		// Generate new message
		msg, tokens, err := llmClient.GenerateCommitMessage(req.Diff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Store in cache
		memCache.Add(req.Diff, msg, tokens)
		
		// Save to history
		if db != nil {
			db.Save(req.Diff, msg, tokens, false)
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   msg,
			"fromCache": false,
			"tokens":    tokens,
		})
	})
	
	// Regenerate endpoint (bypass cache)
	http.HandleFunc("/regenerate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		var req struct {
			Diff string `json:"diff"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		
		// Generate new message (ignore cache)
		msg, tokens, err := llmClient.GenerateCommitMessage(req.Diff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Update cache
		memCache.Add(req.Diff, msg, tokens)
		
		// Save to history
		if db != nil {
			db.Save(req.Diff, msg, tokens, false)
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   msg,
			"fromCache": false,
			"tokens":    tokens,
		})
	})
	
	// History endpoint
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			http.Error(w, "History not available", http.StatusServiceUnavailable)
			return
		}
		
		records, err := db.GetRecent(50)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	})
	
	// Stats endpoint
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"cache_size": memCache.Size(),
			"config": map[string]interface{}{
				"model":        cfg.Model,
				"max_tokens":   cfg.MaxTokens,
				"max_cache":    cfg.MaxCacheSize,
			},
		})
	})
	
	log.Printf("Server starting on http://localhost:%s", cfg.Port)
	log.Printf("Cache size: %d, TTL: 60 minutes", cfg.MaxCacheSize)
	if db != nil {
		log.Println("History database enabled")
	}
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
