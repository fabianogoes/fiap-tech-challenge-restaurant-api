package rest

import "github.com/gin-gonic/gin"

func SetupTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	return r
}
