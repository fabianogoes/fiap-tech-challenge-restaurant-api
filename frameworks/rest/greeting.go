package rest

import "github.com/gin-gonic/gin"

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API GoFood",
	})
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
