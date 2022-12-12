package main

import (
	"fmt"
	"log"
	"net/http"
	"open-notes-api/config"
	"open-notes-api/controllers"
	"open-notes-api/database"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfiguration()
	viper.GetString("PORT")

	database.Connect()
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterRoutes(router)
	router.Use(LogRequest(router))

	log.Printf("Iniciando o servidor na porta %s", viper.GetString("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("PORT")), router))
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/notes/", controllers.GetNotes).Methods("GET", "OPTIONS")
	router.HandleFunc("/notes/{id}/", controllers.GetNoteById).Methods("GET", "OPTIONS")
	router.HandleFunc("/notes/", controllers.CreateNote).Methods("POST", "OPTIONS")
	router.HandleFunc("/notes/{id}/", controllers.UpdateNote).Methods("PUT", "OPTIONS")
	router.HandleFunc("/notes/{id}/", controllers.DeleteNote).Methods("DELETE", "OPTIONS")
}

func LogRequest(next http.Handler) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				log.Printf("%s %s %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path, r.URL.RawQuery)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
