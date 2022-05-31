package core

import (
	"fmt"
	"github.com/vitorsalgado/gopin/internal"
	"github.com/vitorsalgado/gopin/internal/config"
	"github.com/vitorsalgado/gopin/internal/usecases"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/vitorsalgado/gopin/internal/util/panicif"
	"github.com/vitorsalgado/gopin/internal/util/test"
	"github.com/vitorsalgado/gopin/internal/util/worker"
)

var ts *httptest.Server
var repo = FakeRepository{}
var data = []Location{
	{Latitude: 1, Longitude: 1, Precision: 100, ReportedAt: time.Now()},
	{Latitude: 2, Longitude: 2, Precision: 200, ReportedAt: time.Now()},
	{Latitude: 10, Longitude: 12, Precision: 100, ReportedAt: time.Now()},
}

func TestMain(m *testing.M) {
	// Setup and teardown
	srv, r := gopin.Server(config.Load())

	dispatcher := worker.NewDispatcher(2)
	dispatcher.Run()

	usecases.RegisterLocationRoutes(r, dispatcher, &repo)
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
	var result Location
	repo.On("GetCurrent", id).Return(&Location{SessionID: "1000", Latitude: 100, Longitude: 150, Precision: 1500, ReportedAt: time.Now()})

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.GreaterOrEqual(t, result.Precision, 1000.0)
	repo.AssertExpectations(t)
}

func TestItShouldReturnBadRequest_whenParameterIsNotValidUUID(t *testing.T) {
	var id = "test01"
	var result Location

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestItShouldReturnNotFound_whenUnableToRetrieveCurrentLocation(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c2"
	var result Location
	a := &Location{}
	a = nil
	repo.On("GetCurrent", id).Return(a)

	resp := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnLocationHistoryForSession_whenAvailable(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c3"
	var result []Location
	repo.On("HistoryForSession", id).Return(data[:len(data)-1])

	test.GetJSON(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id), &result)

	assert.Equal(t, 2, len(result))
}

func TestItShouldReturn404_whenSessionHistoryIsEmpty(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c4"
	repo.On("HistoryForSession", id).Return([]Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))
	panicif.Err(err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnBadRequest_whenIdIsNotUUID(t *testing.T) {
	id := "test03"
	repo.On("HistoryForSession", id).Return([]Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))
	panicif.Err(err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// Mocks --

type FakeRepository struct {
	mock.Mock
}

func (m *FakeRepository) ReportNew(location Location) {
	m.Called(location)
}

func (m *FakeRepository) GetCurrent(id string) *Location {
	args := m.Called(id)
	return args.Get(0).(*Location)
}

func (m *FakeRepository) HistoryForSession(sessionID string) []Location {
	args := m.Called(sessionID)
	return args.Get(0).([]Location)
}
