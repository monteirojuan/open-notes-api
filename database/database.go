package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	Instance.AutoMigrate(&Note{})
}
