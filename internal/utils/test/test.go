package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vitorsalgado/go-location-management/internal/utils/panicif"
)

// GetJSON is a utility method to make GET HTTP calls with JSON bodied response easier.
// It returns the original http.Response but with Body already closed
func GetJSON(url string, result interface{}) *http.Response {
	resp, err := http.Get(url)
	panicif.Err(err)

	defer func() {
		err = resp.Body.Close()
		panicif.Err(err)
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	panicif.Err(err)

	return resp
}

func PostJSON(url string, content interface{}, result *interface{}) *http.Response {
	data, err := json.Marshal(content)
	panicif.Err(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	panicif.Err(err)

	defer func() {
		err = resp.Body.Close()
		panicif.Err(err)
	}()

	if result != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &result)
		panicif.Err(err)
	}

	return resp
}

// Get simple makes a GET HTTP call without additional behavior
func Get(url string) (*http.Response, error) {
	return http.Get(url)
}
