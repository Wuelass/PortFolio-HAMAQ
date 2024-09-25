package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// database.InitDatabase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

	})

	// Start the server
	router.Run(":8080")

}
