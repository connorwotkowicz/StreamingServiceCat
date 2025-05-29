// This file will: 
// Define what a Media item looks like (struct)
// Set up functions to handle database operations (CRUD)



// Explanation:
// Media is the core data structure — it maps to a row in the media table in your database.
// GetAllMedia is the first DB function we’re writing. It performs a simple SELECT query and returns a list of Media structs.

package models

import (
	"database/sql"

)

// Media represents a movie or show in the catalog
type Media struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
}

// GetAllMedia returns all media records from the database
func GetAllMedia(db *sql.DB) ([]Media, error) {
	rows, err := db.Query("SELECT id, title, description, genre, year FROM media")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaList []Media
	for rows.Next() {
		var m Media
		if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.Genre, &m.Year); err != nil {
			return nil, err
		}
		mediaList = append(mediaList, m)
	}
	return mediaList, nil
}
