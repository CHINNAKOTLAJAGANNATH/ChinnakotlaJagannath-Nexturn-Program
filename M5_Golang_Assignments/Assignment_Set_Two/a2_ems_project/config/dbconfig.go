package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() *sql.DB {
	var err error
	// Open a database connection
	DB, err := sql.Open("sqlite", "./myecommerce.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		stock INTEGER,
		category_id INTEGER
	);`)

	if err != nil {
		log.Fatal("Error creating products table: ", err)
	}

	// Create users tables if they do not exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);`)

	if err != nil {
		log.Fatal("Error creating users table: ", err)
	}

	log.Println("Successfully connected to the product database and ensured the schema exists.")
	return DB

}
