package db

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tree struct {
	gorm.Model
	Name          string
	NameLatin     string
	TreeType      string
	LifeTime      int
	Environnement string
}

type User struct {
	gorm.Model
	Username   string
	Email      string
	Password   string
	Admin      bool
	SessionID  string
	Expiration time.Time
}

func InitDatabase() {
	DB, err := gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}

	if err := DB.AutoMigrate(&User{}, &Tree{}); err != nil {
		log.Fatalln("failed to migrate database: ", err)
	}

	log.Println("Database connection etablished and migratiopn ran successfully.")

}
