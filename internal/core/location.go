package core

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

	// LocationRepository represents data access operations for Location core
	LocationRepository interface {
		ReportNew(location Location)
		GetCurrent(id string) *Location
		HistoryForSession(sessionID string) []Location
	}
)
