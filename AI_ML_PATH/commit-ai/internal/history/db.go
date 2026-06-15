package history

import (
	"database/sql"
	"fmt"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	ID        int
	Diff      string
	Message   string
	Tokens    int
	FromCache bool
	CreatedAt time.Time
}

type DB struct {
	conn *sql.DB
}

// New creates a new history database
func New(dbPath string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	
	// Create table if not exists
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS generations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			diff TEXT NOT NULL,
			message TEXT NOT NULL,
			tokens INTEGER,
			from_cache BOOLEAN,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("create table: %w", err)
	}
	
	// Create index for faster lookups
	_, err = conn.Exec(`
		CREATE INDEX IF NOT EXISTS idx_diff ON generations(diff)
	`)
	if err != nil {
		return nil, fmt.Errorf("create index: %w", err)
	}
	
	return &DB{conn: conn}, nil
}

// Save stores a generation record
func (db *DB) Save(diff, message string, tokens int, fromCache bool) error {
	_, err := db.conn.Exec(
		"INSERT INTO generations (diff, message, tokens, from_cache) VALUES (?, ?, ?, ?)",
		diff, message, tokens, fromCache,
	)
	if err != nil {
		return fmt.Errorf("save record: %w", err)
	}
	return nil
}

// GetRecent returns the last N records
func (db *DB) GetRecent(limit int) ([]Record, error) {
	rows, err := db.conn.Query(
		"SELECT id, diff, message, tokens, from_cache, created_at FROM generations ORDER BY created_at DESC LIMIT ?",
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("query recent: %w", err)
	}
	defer rows.Close()
	
	var records []Record
	for rows.Next() {
		var r Record
		err := rows.Scan(&r.ID, &r.Diff, &r.Message, &r.Tokens, &r.FromCache, &r.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		records = append(records, r)
	}
	
	return records, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}
