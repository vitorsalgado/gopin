package current

import (
	"github.com/vitorsalgado/go-location-management/internal/locations/domain"
)

type (
	// Location represents the current location use case
	Location struct {
		repository domain.Repository
	}
)

// NewCurrentLocation creates current location use case instance
func NewCurrentLocation(repository domain.Repository) *Location {
	return &Location{repository}
}
