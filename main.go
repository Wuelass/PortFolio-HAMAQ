package main

import (
	"PortFolio-HAMAQ/db"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDatabase()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/register", func(c *gin.Context) {

		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirmPassword := c.PostForm("confirmPassword")
		password, err := db.HashPassword(password)
		if err != nil {
			log.Fatalln("error hashing password", err)
		}

		if !db.CheckPasswordHash(confirmPassword, password) {
			log.Fatalln("password and confirmPassword are not the same", err)
		}

		err = db.AddUser(username, email, password)
		if err != nil {
			log.Fatalln("error while adding user", err)
		}

	})

	router.POST("/login", func(c *gin.Context) {
		var connected bool
		if !connected {
			var user db.User
			var userFound bool
			var err error
			var connectionApprouved bool
			identifiant := c.PostForm("identifiant")
			password := c.PostForm("password")
			if identifiant != "" {
				user, err = db.GetUserByEmail(identifiant)
				if err == nil {
					userFound = true
				}
				if !userFound {
					user, err = db.GetUserByUsername(identifiant)
					if err == nil {
						userFound = true
					} else {
						log.Fatalln("no user found")
					}
				}
			}

			if userFound {
				connectionApprouved = db.CheckPasswordHash(password, user.Password)
			}

			if connectionApprouved {
				connected = true
			}
		}

	})

	router.GET("/arbre", func(c *gin.Context) {

	})

	router.GET("/arbre/:id", func(c *gin.Context) {
		idStr := c.Param("id") // Extract the ID as a string.

		// Convert the ID from string to int.
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Call GetTreeById to fetch the tree.
		tree, err := db.GetTreeById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tree not found"})
			return
		}

		// Return the tree data as JSON.
		c.JSON(http.StatusOK, tree)
	})

	router.POST("/arbre", func(c *gin.Context) {
		name := c.PostForm("name")
		name_latin := c.PostForm("name_latin")
		race := c.PostForm("race")
		lifetime := c.PostForm("lifetime")
		biome := c.PostForm("biome")

		err := db.AddTree(name, name_latin, race, lifetime, biome)
		if err != nil {
			log.Fatalln("error while adding Tree", err)
		}
	})

	router.PUT("/arbre/:id", func(c *gin.Context) {

	})

	router.DELETE("/arbre/:id", func(c *gin.Context) {
		idStr := c.Param("id") // Extract the ID from the URL

		// Convert the ID from string to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.CheckTreeByID(id); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Tree not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		if err := db.DeleteTreeByID(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tree"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tree deleted successfully"})
	})

	// Start the server
	router.Run(":8080")

}
