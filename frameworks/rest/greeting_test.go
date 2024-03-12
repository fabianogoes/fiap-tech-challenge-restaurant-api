package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcome(t *testing.T) {
	r := SetupTest()
	r.GET("/", Welcome)
	request, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

	var welcomeResponse struct {
		Message string
	}
	err = json.Unmarshal(response.Body.Bytes(), &welcomeResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Welcome to the API GoFood", welcomeResponse.Message)
}

func TestHealth(t *testing.T) {
	r := SetupTest()
	r.GET("/health", Health)
	request, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

	var healthResponse struct {
		Status string
	}
	err = json.Unmarshal(response.Body.Bytes(), &healthResponse)
	assert.NoError(t, err)
	assert.Equal(t, "UP", healthResponse.Status)
}
