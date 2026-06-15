package cache

import (
	"sync"
	"time"
)

type Entry struct {
	Diff      string
	Message   string
	Tokens    int
	CreatedAt time.Time
}

type Cache struct {
	entries   []Entry
	maxSize   int
	ttl       time.Duration
	mu        sync.RWMutex
}

// New creates a new cache with max size and TTL
func New(maxSize int, ttlMinutes int) *Cache {
	return &Cache{
		entries: make([]Entry, 0, maxSize),
		maxSize: maxSize,
		ttl:     time.Duration(ttlMinutes) * time.Minute,
	}
}

// Add adds or updates an entry (prepends, keeps maxSize)
func (c *Cache) Add(diff, message string, tokens int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	// Remove existing entry with same diff if present
	for i, entry := range c.entries {
		if entry.Diff == diff {
			c.entries = append(c.entries[:i], c.entries[i+1:]...)
			break
		}
	}
	
	// Add new entry at front
	entry := Entry{
		Diff:      diff,
		Message:   message,
		Tokens:    tokens,
		CreatedAt: time.Now(),
	}
	c.entries = append([]Entry{entry}, c.entries...)
	
	// Trim to max size
	if len(c.entries) > c.maxSize {
		c.entries = c.entries[:c.maxSize]
	}
}

// Get retrieves an entry if it exists and hasn't expired
func (c *Cache) Get(diff string) (string, int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	for _, entry := range c.entries {
		if entry.Diff == diff {
			// Check if expired
			if time.Since(entry.CreatedAt) > c.ttl {
				return "", 0, false
			}
			return entry.Message, entry.Tokens, true
		}
	}
	return "", 0, false
}

// Size returns current cache size
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}
