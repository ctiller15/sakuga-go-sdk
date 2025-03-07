package sakugaapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIInitialization(t *testing.T) {
	newAPI := NewAPI()

	assert.NotNil(t, newAPI)
	assert.NotNil(t, newAPI.Posts)
}
