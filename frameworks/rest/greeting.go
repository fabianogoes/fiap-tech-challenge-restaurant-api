package rest

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API GoFood",
	})
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func Environment(c *gin.Context) {
	config := entities.NewConfig()

	c.JSON(http.StatusOK, config)
}
