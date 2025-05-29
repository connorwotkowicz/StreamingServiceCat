// This file will:

// Use database/sql to connect to your DB
// Load credentials from environment variables
// Expose a shared DB instance to the app


// Explanation:
// Uses environment variables to avoid hardcoding secrets.
// Exposes a global DB variable you can import elsewhere.
// sslmode=disable is fine for local dev. You’d use require or verify-full in production.


package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB connection:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	log.Println("Database connection established")
}
