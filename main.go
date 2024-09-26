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

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.POST("/registerUser", func(c *gin.Context) {
		var err error
		var hashedPassword string
		username := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirmPassword := c.PostForm("confirmPassword")
		if password == confirmPassword {
			hashedPassword, err = db.HashPassword(password)
			if err != nil {
				log.Fatalln("error hashing password", err)
			}
			log.Println("Hashed Password:", hashedPassword)
			if err := db.AddUser(username, email, hashedPassword); err != nil {
				// Log the error (for internal debugging purposes)
				log.Printf("error while adding user: %v", err)
				// Redirect back to the registration page with an error message
				c.HTML(http.StatusOK, "register.html", gin.H{
					"error":    err.Error(), // Pass the error message to the template
					"username": username,    // Preserve entered values so the user doesn't need to re-enter them
					"email":    email,
				})

				return
			}
		} else {
			log.Println("password and confirmPassword are not the same")
			c.Redirect(http.StatusSeeOther, "/register")
			return
		}

		c.Redirect(http.StatusSeeOther, "/login")

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
		var trees []db.Tree

		if err := db.DB.Find(&trees).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trees"})
			return
		}

		c.HTML(http.StatusOK, "alltree.html", trees)
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

	router.GET("/add-tree", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_tree.html", nil)
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

	router.POST("/arbresupp", func(c *gin.Context) {
		idStr := c.PostForm("id")
		log.Println("Received ID:", idStr) // Log the received ID

		// Convert the ID from string to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Invalid ID format:", idStr) // Log invalid ID
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		log.Println("Converted ID:", id) // Log the converted ID

		// Check if the tree with the given ID exists
		if err := db.CheckTreeByID(id); err != nil {
			log.Println("Error checking tree by ID:", err) // Log the error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Tree not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		// Proceed to delete the tree
		if err := db.DeleteTreeByID(id); err != nil {
			log.Println("Error deleting tree by ID:", err) // Log the error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tree"})
			return
		}

		log.Println("Tree deleted successfully with ID:", id) // Log success
		c.HTML(http.StatusOK, "successSupp.html", nil)
	})

	// Start the server
	router.Run(":8080")

}
