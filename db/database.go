package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Admin    bool
}

func InitDatabase() {
	DB, err := gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalln("failed to migrate database: ", err)
	}

	log.Println("Database connection etablished and migratiopn ran successfully.")

}
