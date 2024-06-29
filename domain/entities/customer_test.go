package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Customer(t *testing.T) {
	customer, err := NewCustomer("test", "test@test.com", "123")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}
