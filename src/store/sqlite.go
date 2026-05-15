package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db            *sql.DB
	latestContent string
}

func NewSQLiteStore(dbPath string) *SQLiteStore {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS clipboard (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	store := &SQLiteStore{
		db: db,
	}

	// Load latest content
	store.loadLatestContent()

	return store
}

func (s *SQLiteStore) loadLatestContent() {
	var content string
	err := s.db.QueryRow("SELECT content FROM clipboard ORDER BY timestamp DESC LIMIT 1").Scan(&content)
	if err == nil {
		s.latestContent = content
	}
}

func (s *SQLiteStore) GetLatestContent() string {
	return s.latestContent
}

func (s *SQLiteStore) SetLatestContent(content string) {
	if content != s.latestContent {
		_, err := s.db.Exec("INSERT INTO clipboard (content) VALUES (?)", content)
		if err != nil {
			log.Printf("Error storing clipboard content: %v", err)
			return
		}
		s.latestContent = content
	}
}

func (s *SQLiteStore) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
