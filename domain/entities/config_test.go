package entities

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ConfigDefault(t *testing.T) {
	config := NewConfig()
	assert.NotNil(t, config)
	assert.Equal(t, "default", config.Environment)
}

func Test_ConfigDevelopment(t *testing.T) {
	err := os.Setenv("APP_ENV", "development")
	config := NewConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "development", config.Environment)
}

func Test_ConfigProduction(t *testing.T) {
	err := os.Setenv("APP_ENV", "production")
	config := NewConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "production", config.Environment)
}
