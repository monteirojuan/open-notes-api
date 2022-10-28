package controllers

import (
    "encoding/json"
    "net/http"
    "open-notes-api/database"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
    var note []database.Note

    database.Instance.Find(&note)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(note)
}
