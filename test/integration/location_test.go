package integration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vitorsalgado/gopin/internal"
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/util/config"
	"github.com/vitorsalgado/gopin/internal/util/test"
	"github.com/vitorsalgado/gopin/internal/util/worker"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
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
	srv, r := gopin.Server(config.Load())

	dispatcher := worker.NewDispatcher(2)
	dispatcher.Run()

	gopin.Routes(r, dispatcher, &repo)
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
	repo.On("Current", id).Return(&domain.Location{SessionID: "1000", Latitude: 100, Longitude: 150, Precision: 1500, ReportedAt: time.Now()})

	resp, err := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.GreaterOrEqual(t, result.Precision, 1000.0)
	repo.AssertExpectations(t)
}

func TestItShouldReturnBadRequest_whenParameterIsNotValidUUID(t *testing.T) {
	var id = "test01"
	var result domain.Location

	resp, err := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestItShouldReturnNotFound_whenUnableToRetrieveCurrentLocation(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c2"
	var result domain.Location
	a := &domain.Location{}
	a = nil
	repo.On("Current", id).Return(a)

	resp, err := test.GetJSON(fmt.Sprintf("%s/api/v1/current_location/%v", ts.URL, id), &result)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnLocationHistoryForSession_whenAvailable(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c3"
	var result []domain.Location
	repo.On("HistoryForSession", id).Return(data[:len(data)-1])

	_, err := test.GetJSON(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id), &result)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func TestItShouldReturn404_whenSessionHistoryIsEmpty(t *testing.T) {
	var id = "79561481-fc11-419c-a9e8-e5a079b853c4"
	repo.On("HistoryForSession", id).Return([]domain.Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestItShouldReturnBadRequest_whenIdIsNotUUID(t *testing.T) {
	id := "test03"
	repo.On("HistoryForSession", id).Return([]domain.Location{})

	resp, err := test.Get(fmt.Sprintf("%s/api/v1/location_history/%v", ts.URL, id))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// Mocks --

type FakeRepository struct {
	mock.Mock
}

func (m *FakeRepository) ReportNew(location domain.Location) error {
	return m.Called(location).Get(0).(error)
}

func (m *FakeRepository) Current(id string) (*domain.Location, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Location), args.Get(1).(error)
}

func (m *FakeRepository) HistoryForSession(sessionID string) (*[]domain.Location, error) {
	args := m.Called(sessionID)
	return args.Get(0).(*[]domain.Location), args.Get(1).(error)
}
