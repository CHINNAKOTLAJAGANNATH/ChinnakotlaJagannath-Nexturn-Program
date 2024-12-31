package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

// InitializeDatabase initializes the SQLite database and ensures the schema exists
func InitializeDatabase() *sql.DB {
	var err error
	// Open the SQLite database file
	DB, err = sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the blogs table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		author TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the users table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the blog database and ensured the schema exists.")
	return DB
}
