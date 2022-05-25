package locations

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/vitorsalgado/gopin/internal/locations/domain"
	"github.com/vitorsalgado/gopin/internal/server/rest"
	"github.com/vitorsalgado/gopin/internal/utils/config"
	"github.com/vitorsalgado/gopin/internal/utils/panicif"
	"github.com/vitorsalgado/gopin/internal/utils/test"
	"github.com/vitorsalgado/gopin/internal/utils/worker"
)

var ts *httptest.Server
var repo = FakeRepository{}
var data = []domain.Location{
	{Latitude: 1, Longitude: 1, Precision: 100, ReportedAt: time.Now()},
	{Latitude: 2, Longitude: 2, Precision: 200, ReportedAt: time.Now()},
	{Latitude: 10, Longitude: 12, Precision: 100, ReportedAt: time.Now()},
}

func TestMain(m *testing.M) {
	// Setup and teardown
	srv, r := rest.Server(config.Load())

	dispatcher := worker.NewDispatcher(2)
	dispatcher.Run()

	RegisterRoutes(r, dispatcher, &repo)
	r.ApplyRoutesTo(srv)

	ts = httptest.NewServer(srv)
	defer ts.Close()

	// Actual test execution
	code := m.Run()

	// Exit runner with test execution status code
	os.Exit(code)
}

func TestItShouldReturnTheCurrentLocation(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c1"
	var result domain.Location
	repo.On("GetCurrentLocation", id).Return(&domain.Location{SessionID: "1000", Latitude: 100, Longitude: 150, Precision: 1500, ReportedAt: time.Now()})

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.GreaterOrEqual(t, result.Precision, 1000.0)
	repo.AssertExpectations(t)
}

func TestItShouldReturnBadRequest_whenParameterIsNotValidUUID(t *testing.T) {
	var id = "test01"
	var result domain.Location

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestItShouldReturnNotFound_whenUnableToRetrieveCurrentLocation(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c2"
	var result domain.Location
	a := &domain.Location{}
	a = nil
	repo.On("GetCurrentLocation", id).Return(a)

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnLocationHistoryForSession_whenAvailable(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c3"
	var result []domain.Location
	repo.On("HistoryFor", id).Return(data[:len(data)-1])

	test.GetJSON(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id), &result)

	assert.Equal(t, 2, len(result))
}

func TestItShouldReturn404_whenSessionHistoryIsEmpty(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c4"
	repo.On("HistoryFor", id).Return([]domain.Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))
	panicif.Err(err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnBadRequest_whenIdIsNotUUID(t *testing.T) {
	id := "test03"
	repo.On("HistoryFor", id).Return([]domain.Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))
	panicif.Err(err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// Mocks --

type FakeRepository struct {
	mock.Mock
}

func (m *FakeRepository) ReportNewLocation(location domain.Location) {
	m.Called(location)
}

func (m *FakeRepository) GetCurrentLocation(id string) *domain.Location {
	args := m.Called(id)
	return args.Get(0).(*domain.Location)
}

func (m *FakeRepository) HistoryFor(sessionID string) []domain.Location {
	args := m.Called(sessionID)
	return args.Get(0).([]domain.Location)
}
