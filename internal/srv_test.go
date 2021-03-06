package gopin

import (
	"fmt"
	"github.com/vitorsalgado/gopin/internal/util/config"
	"github.com/vitorsalgado/gopin/internal/util/observability"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/gopin/internal/util/test"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	// Setup and teardown
	var conf config.Config
	conf.SwaggerUiPath = "../docs/openapi/swagger-ui"

	srv, r := Server(&conf)
	observability.ConfigureHealthCheck(r)
	r.ApplyRoutesTo(srv)

	ts = httptest.NewServer(srv)
	defer ts.Close()

	// Actual test execution
	code := m.Run()

	// Exit runner with test execution status code
	os.Exit(code)
}

func TestItShouldReturnPongWhenServeIsLive(t *testing.T) {
	var result observability.Result
	res, err := test.GetJSON(fmt.Sprintf("%s/api/v1/ping", ts.URL), &result)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "pong", result.Status)
}

func TestItShouldReturnSwaggerDocs(t *testing.T) {
	resp, err := test.Get(fmt.Sprintf("%s/docs/", ts.URL))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "text/html; charset=utf-8", resp.Header.Get("Content-Type"))
}
