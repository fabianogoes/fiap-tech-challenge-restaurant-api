package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setup() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

func TestWelcome(t *testing.T) {
	r := setup()
	r.GET("/", Welcome)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestHealth(t *testing.T) {
	r := setup()
	r.GET("/health", Health)
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
