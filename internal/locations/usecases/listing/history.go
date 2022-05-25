package listing

import (
	"github.com/vitorsalgado/gopin/internal/locations/domain"
)

type (
	// History represents the session location history listing use case
	History struct {
		repository domain.Repository
	}
)

// NewHistory creates a new History listing use case instance
func NewHistory(repository domain.Repository) *History {
	return &History{repository}
}
