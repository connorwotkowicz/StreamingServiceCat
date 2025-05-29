// What this does:
// Calls models.GetAllMedia(db.DB) to get the list from your DB
// If it fails, returns a 500 error
// Otherwise, sends a JSON response with all the media records

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"streamflix-api/db"
	"streamflix-api/models"

	"github.com/go-chi/chi/v5"
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


func GetMediaByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	media, err := models.GetMediaByID(db.DB, id)
	if err != nil {
		http.Error(w, "Failed to fetch media", http.StatusInternalServerError)
		return
	}
	if media == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(media)
}






func CreateMedia(w http.ResponseWriter, r *http.Request) {
	var m models.Media
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := models.CreateMedia(db.DB, m)
	if err != nil {
		http.Error(w, "Failed to create media", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}



func UpdateMedia(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var m models.Media
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = models.UpdateMedia(db.DB, id, m)
	if err != nil {
		http.Error(w, "Failed to update media", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}





func DeleteMedia(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Printf("Attempting to delete media with ID: %d\n", id)

	err = models.DeleteMedia(db.DB, id)
	if err != nil {
		log.Printf("Delete failed: %v\n", err)
		http.Error(w, "Failed to delete media", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully deleted media with ID: %d\n", id)
	w.WriteHeader(http.StatusNoContent)
}
