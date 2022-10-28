package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DATABASE_HOST"),
		viper.GetString("DATABASE_USER"),
		viper.GetString("DATABASE_PASSWORD"),
		viper.GetString("DATABASE_NAME"),
		viper.GetString("DATABASE_SSL"),
	)
	log.Println(dsn)

	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	Instance.AutoMigrate(&Note{})
}
