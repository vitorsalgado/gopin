package checks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldReturnTrue_whenUUIDIsValid(t *testing.T) {
	assert.True(t, IsUUIDValid("79561481-fc11-419c-a9e8-e5a079b853c5"))
}

func TestItShouldReturnFalse_whenUUIDIsNotValid(t *testing.T) {
	assert.False(t, IsUUIDValid("a9e8-e5a079b853c5"))
}
