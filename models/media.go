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

func GetMediaByID(db *sql.DB, id int) (*Media, error) {
	query := "SELECT id, title, description, genre, year FROM media WHERE id = $1"

	var m Media
	err := db.QueryRow(query, id).Scan(&m.ID, &m.Title, &m.Description, &m.Genre, &m.Year)
	if err == sql.ErrNoRows {
		return nil, nil // not found
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}



func CreateMedia(db *sql.DB, m Media) (int, error) {
	query := `
		INSERT INTO media (title, description, genre, year)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int
	err := db.QueryRow(query, m.Title, m.Description, m.Genre, m.Year).Scan(&id)
	return id, err
}



func UpdateMedia(db *sql.DB, id int, m Media) error {
	query := `
		UPDATE media
		SET title = $1, description = $2, genre = $3, year = $4
		WHERE id = $5
	`
	_, err := db.Exec(query, m.Title, m.Description, m.Genre, m.Year, id)
	return err
}




func DeleteMedia(db *sql.DB, id int) error {
	query := "DELETE FROM media WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}
