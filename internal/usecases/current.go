package usecases

import (
	"github.com/vitorsalgado/gopin/internal/core"
)

type (
	// Location represents the current location use case
	Location struct {
		repository core.LocationRepository
	}
)

// NewCurrentLocation creates current location use case instance
func NewCurrentLocation(repository core.LocationRepository) *Location {
	return &Location{repository}
}
