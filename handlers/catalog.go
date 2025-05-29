
// What this does:
// Calls models.GetAllMedia(db.DB) to get the list from your DB
// If it fails, returns a 500 error
// Otherwise, sends a JSON response with all the media records

package handlers

import (
	"encoding/json"
	"net/http"

	"streamflix-api/db"
	"streamflix-api/models"
)

func GetAllMedia(w http.ResponseWriter, r *http.Request) {
	mediaList, err := models.GetAllMedia(db.DB)
	if err != nil {
		http.Error(w, "Failed to fetch media", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mediaList)
}
