package db

import (
	"log"
	//"time"

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
	//SessionID  string
	//Expiration time.Time
}

var DB *gorm.DB

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

func AddUser(username, email, password string) error {
	user := User{
		Username: username,
		Email : email,
		Password: password,
		Admin: false,
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}


func GetUserByUsername(username string) (User, error){
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return User{}, result.Error // Return an empty User and the error
	}
	return user, nil
}
func GetUserByEmail(email string) (User, error){
	var user User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, result.Error // Return an empty User and the error
	}
	return user, nil
}



