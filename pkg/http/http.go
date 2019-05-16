package http

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Client http client
type Client struct {
	baseURL string
}

// NewClient create a new client
func NewClient() (Client, error) {
	client := Client{}
	return client, nil
}

// BaseURL set base url
func (client *Client) BaseURL(baseURL string) {
	client.baseURL = baseURL
}

// Get make a get request
func (client *Client) Get(url string, params map[string]string) ([]byte, error) {
	if client.baseURL == "" {
		return []byte{}, errors.New("Base URL not set")
	}
	request, err := http.NewRequest("GET", client.baseURL+url, nil)
	if err != nil {
		return []byte{}, nil
	}
	// if you appending to existing query this works fine
	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v)

	}
	request.URL.RawQuery = q.Encode()

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
