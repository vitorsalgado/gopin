package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetJSON is a utility method to make GET HTTP calls with JSON bodied response easier.
// It returns the original http.Response but with Body already closed
func GetJSON(url string, result interface{}) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PostJSON(url string, content interface{}, result *interface{}) (*http.Response, error) {
	data, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if result != nil {
		body, _ := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// Get simple makes a GET HTTP call without additional behavior
func Get(url string) (*http.Response, error) {
	return http.Get(url)
}
