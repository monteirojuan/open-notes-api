package config

import (
	"github.com/spf13/viper"
	"log"
)

// LoadConfiguration carrega as configurações usando Viper
func LoadConfiguration() {
	log.Println("Carregando configurações...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
