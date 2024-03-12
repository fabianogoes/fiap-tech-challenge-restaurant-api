package rest

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcome(t *testing.T) {
	r := SetupTest()
	r.GET("/", Welcome)
	request, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestHealth(t *testing.T) {
	r := SetupTest()
	r.GET("/health", Health)
	request, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
