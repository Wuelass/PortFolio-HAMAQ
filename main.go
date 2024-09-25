package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

	})

	router.GET("/register", func(c *gin.Context) {

	})

	router.GET("/login", func(c *gin.Context) {

	})

	router.GET("/arbre", func(c *gin.Context) {

	})

	router.POST("/arbre", func(c *gin.Context) {

	})

	router.PUT("/arbre", func(c *gin.Context) {

	})

	router.DELETE("/arbre", func(c *gin.Context) {

	})

	// Start the server
	router.Run(":8080")

}
