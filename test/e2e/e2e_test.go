package e2e

import (
	"fmt"
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/handlers"
	"github.com/vitorsalgado/gopin/internal/util/config"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/gopin/internal/util/test"
)

var configurations *config.Config
var baseURL string
var seed Seed

func TestMain(m *testing.M) {
	configurations = config.Load()

	db := ConnectDb(20 * time.Second)
	baseURL = fmt.Sprintf("http://localhost%v", configurations.Port)
	seed = Seed{db}

	seed.cleanDb()
	seed.seed()

	code := m.Run()

	seed.cleanDb()

	os.Exit(code)
}

func TestPing(t *testing.T) {
	resp, err := test.Get(fmt.Sprintf("%v/api/v1/ping", baseURL))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCurrentLocation(t *testing.T) {
	t.Run("it should return the newest and most precise location as the current", func(t *testing.T) {
		var s1 *domain.Location
		resp, err := test.GetJSON(fmt.Sprintf("%s/%v/%v", baseURL, "api/v1/current_location", u1), &s1)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, session1, s1.SessionID)
	})

	t.Run("it should return 404 (Not Found) when the last location update is older than 10 Minutes",
		func(t *testing.T) {
			resp2, err := test.GetJSON(
				fmt.Sprintf("%s/%v/%v", baseURL, "api/v1/current_location", u2), nil)

			assert.Nil(t, err)
			assert.Equal(t, http.StatusNotFound, resp2.StatusCode)
		})
}

func TestSessionHistory(t *testing.T) {
	t.Run("it should return 404 (Not Found) when there is no location history record for session", func(t *testing.T) {
		resp, err := test.GetJSON(fmt.Sprintf("%v/%v/%v", baseURL, "api/v1/session_location_history", "79561481-fc11-419c-a9e8-e5a079b853c5"), nil)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it should return the location history of the session when available", func(t *testing.T) {
		var history []domain.Location
		resp, err := test.GetJSON(fmt.Sprintf("%v/%v/%v", baseURL, "api/v1/session_location_history", session2), &history)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, 2, len(history))
	})
}

func TestReportNewLocation(t *testing.T) {
	var user = "79561481-fc11-419c-a9e8-e5a079b853c5"
	var session = "daff4b9f-24e2-478d-8b42-6d3f59a08b35"

	var url = fmt.Sprintf("%v/api/v1/location", baseURL)

	_, err := test.PostJSON(url, handlers.ReportLocationRequest{
		UserID:     user,
		SessionID:  session,
		Latitude:   -33.22325847832756,
		Longitude:  -70.21369951517998,
		Precision:  1000,
		ReportedAt: time.Now(),
	}, nil)

	assert.Nil(t, err)

	// We wait enough time so the worker pool can process the
	// location update job
	time.Sleep(5 * time.Second)

	var history []domain.Location
	resp, err := test.GetJSON(fmt.Sprintf("%v/%v/%v", baseURL, "api/v1/session_location_history", session), &history)

	// Now we assert that there is 1 location entry for the session we previously added
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1, len(history))
}
