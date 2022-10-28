package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"open-notes-api/config"
    "open-notes-api/controllers"
    "open-notes-api/database"
)

func main() {
	config.LoadConfiguration()

	database.Connect(viper.GetString("database.dsn"))
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterRoutes(router)

	log.Println(fmt.Sprintf("Iniciando o servidor na porta %s", viper.GetString("server.port")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("server.port")), router))
}

func RegisterRoutes(router *mux.Router) {
	// Umidade do solo
	router.HandleFunc("api/umidade", controllers.GetNote).Methods("GET")
}
