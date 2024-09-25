package main

import (
	"PortFolio-HAMAQ/db"
	"log"

	"github.com/gin-gonic/gin"

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

	router.PUT("/arbre", func(c *gin.Context) {

	})

	router.DELETE("/arbre", func(c *gin.Context) {

	})

	// Start the server
	router.Run(":8080")

}
