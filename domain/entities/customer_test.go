package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Customer(t *testing.T) {
	customer := NewCustomer("test", "test@test.com", "123")
	assert.NotNil(t, customer)
}
