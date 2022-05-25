package panicif

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldPanicIfErrorIsNotNil(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	Err(errors.New("failure"))
}

func TestItShouldNotPanicIfErrorIsNil(t *testing.T) {
	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	Err(nil)
}
