package database

import (
	"fmt"

	"github.com/Pebody-h/Anime-Api/models"
	"github.com/Pebody-h/Anime-Api/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbHost := settings.Settings("DB_HOST")
	dbPort := settings.Settings("DB_PORT")
	dbUser := settings.Settings("DB_USER")
	dbPassword := settings.Settings("DB_PASSWORD")
	dbName := settings.Settings("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail connection to database")
	}

	DB.AutoMigrate(&models.Animes{})
}
