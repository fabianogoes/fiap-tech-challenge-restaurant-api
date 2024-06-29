package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Attendant(t *testing.T) {
	attendant, err := NewAttendant("test")
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}
