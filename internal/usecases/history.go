package usecases

import (
	"github.com/vitorsalgado/gopin/internal/core"
)

type (
	// History represents the session location history listing use case
	History struct {
		repository core.LocationRepository
	}
)

// NewHistory creates a new History listing use case instance
func NewHistory(repository core.LocationRepository) *History {
	return &History{repository}
}
