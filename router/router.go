package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router
	router := gin.Default()
	//
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Run the server
	router.Run(":8080")
}
