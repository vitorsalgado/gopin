package rest

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/gopin/internal/utils/config"
	"github.com/vitorsalgado/gopin/internal/utils/test"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	// Setup and teardown
	var conf config.Config
	conf.SwaggerUiPath = "../../../api/swagger-ui"

	srv, r := Server(&conf)
	RegisterRoutes(r)
	r.ApplyRoutesTo(srv)

	ts = httptest.NewServer(srv)
	defer ts.Close()

	// Actual test execution
	code := m.Run()

	// Exit runner with test execution status code
	os.Exit(code)
}

func TestItShouldReturnPongWhenServeIsLive(t *testing.T) {
	var result result
	response := test.GetJSON(fmt.Sprintf("%s/api/v1/ping", ts.URL), &result)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "pong", result.Status)
}

//func TestItShouldReturnSwaggerDocs(t *testing.T) {
//	resp, err := test.Get(fmt.Sprintf("%s/docs/", ts.URL))
//	panicif.Err(err)
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//	assert.Equal(t, "text/html; charset=utf-8", resp.Header.Get("Content-Type"))
//}
