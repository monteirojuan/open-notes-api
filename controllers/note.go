package controllers

import (
	"encoding/json"
	"net/http"
	"open-notes-api/database"

	"github.com/gorilla/mux"
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

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	note_id := mux.Vars(r)["id"]
	var note database.Note
	database.Instance.First(&note, note_id)

	if note.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	note_id := mux.Vars(r)["id"]
	var note database.Note
	database.Instance.First(&note, note_id)

	if note.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&note)
	database.Instance.Save(&note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	note_id := mux.Vars(r)["id"]
	var note database.Note
	database.Instance.First(&note, note_id)

	if note.ID == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	database.Instance.Delete(&note, note_id)
	w.WriteHeader(http.StatusOK)
}
