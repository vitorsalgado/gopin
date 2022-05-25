package domain

import (
	"time"
)

type (
	Location struct {
		UserID     string    `json:"user_uuid,omitempty"`
		SessionID  string    `json:"session_uuid,omitempty"`
		Latitude   float64   `json:"lat,omitempty"`
		Longitude  float64   `json:"lng,omitempty"`
		Precision  float64   `json:"precision,omitempty"`
		ReportedAt time.Time `json:"reported_at,omitempty"`
	}

	// Repository represents data access operations for Location domain
	Repository interface {
		ReportNewLocation(location Location)
		GetCurrentLocation(id string) *Location
		HistoryFor(sessionID string) []Location
	}
)
