package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfiguration carrega as configurações usando Viper
func LoadConfiguration() {
	log.Println("Carregando configurações...")

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("DATABASE_HOST", "localhost")
	viper.SetDefault("DATABASE_USER", "postgres")
	viper.SetDefault("DATABASE_PASSWORD", "password")
	viper.SetDefault("DATABASE_NAME", "postgres")
	viper.SetDefault("DATABASE_SSL", "disable")
}
