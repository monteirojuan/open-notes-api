package controllers

import (
	"encoding/json"
	"net/http"
	"open-notes-api/database"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note database.Note
	json.NewDecoder(r.Body).Decode(&note)
	database.Instance.Create(&note)
	json.NewEncoder(w).Encode(note)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var note []database.Note

	database.Instance.Find(&note)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}
