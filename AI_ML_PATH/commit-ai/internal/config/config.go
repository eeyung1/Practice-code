package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey      string `json:"api_key"`
	Model       string `json:"model"`
	MaxTokens   int    `json:"max_tokens"`
	MaxCacheSize int   `json:"max_cache_size"`
	Port        string `json:"port"`
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Model:        "llama-3.1-8b-instant",
		MaxTokens:    200,
		MaxCacheSize: 10,
		Port:         "8080",
	}
}

// Load reads config from file or environment
func Load(configPath string) (*Config, error) {
	cfg := DefaultConfig()
	
	// Try to read from file
	if configPath != "" {
		data, err := os.ReadFile(configPath)
		if err == nil {
			if err := json.Unmarshal(data, cfg); err != nil {
				return nil, fmt.Errorf("invalid config: %w", err)
			}
		}
	}
	
	// Environment variables override file
	if apiKey := os.Getenv("GROQ_API_KEY"); apiKey != "" {
		cfg.APIKey = apiKey
	}
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY not set (either in config file or environment)")
	}
	
	return cfg, nil
}

// Save writes config to file
func (c *Config) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}
