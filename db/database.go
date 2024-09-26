package db

import (
	"errors"
	"fmt"
	"log"

	//"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TreeList struct {
	Trees []Tree
}

type Tree struct {
	gorm.Model
	Name          string
	NameLatin     string
	TreeType      string
	LifeTime      string
	Environnement string
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Admin    bool
	//SessionID  string
	//Expiration time.Time
}

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}

	if err := DB.AutoMigrate(&User{}, &Tree{}); err != nil {
		log.Fatalln("failed to migrate database: ", err)
	}

	log.Println("Database connection etablished and migratiopn ran successfully.")

}

func AddUser(username, email, password string) error {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err == nil {
		return fmt.Errorf("username already taken")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := DB.Where("email = ?", email).First(&user).Error; err == nil {
		return fmt.Errorf("email already taken")
	} else if !errors.Is(err, gorm.ErrRecordNotFound){
		return err
	}

	newUser := User{
		Username: username,
		Email:    email,
		Password: password,
		Admin:    false,
	}
	if err := DB.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return User{}, result.Error // Return an empty User and the error
	}
	return user, nil
}

func GetTreeById(id int) (Tree, error) {
	var tree Tree
	result := DB.Where("id = ?", id).First(&tree)
	if result.Error != nil {
		return Tree{}, result.Error
	}
	return tree, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, result.Error // Return an empty User and the error
	}
	return user, nil
}

func AddTree(name, nameLatin, treeType, lifetime string, environnement string) error {
	if DB == nil {
		log.Println("Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}

	tree := Tree{
		Name:          name,
		NameLatin:     nameLatin,
		TreeType:      treeType,
		LifeTime:      lifetime,
		Environnement: environnement,
	}
	if err := DB.Create(&tree).Error; err != nil {
		return err
	}
	log.Println("tree add")
	return nil

}

func DeleteTreeByID(id int) error {
	if err := DB.Where("id = ?", id).Delete(&Tree{}).Error; err != nil {
		return err
	}
	return nil
}

func CheckTreeByID(id int) error {
	var tree Tree
	// Query the database to check if a tree with the given ID exists
	if err := DB.First(&tree, id).Error; err != nil {
		return err // Return the error (e.g., "record not found" or database errors)
	}
	return nil // If the tree exists, return nil (no error)
}
