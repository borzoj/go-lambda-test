package weather

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/borzoj/go-lambda-test/pkg/http"
)

// Service weather api service
type Service struct {
	client http.Client
	appID  string
}

// Weather description
type Weather struct {
	Description string `json:"description" binding:"required"`
	Main        string `json:"main" binding:"required"`
}

// Response api description
type Response struct {
	Weather []Weather `json:"weather" binding:"required"`
}

// NewService return a new servide
func NewService(client http.Client) (Service, error) {
	return Service{client: client}, nil
}

//AppID set app id for authenticating requests
func (service *Service) AppID(appID string) {
	service.appID = appID
}

// BaseURL set base url for request
func (service *Service) BaseURL(baseURL string) {
	service.client.BaseURL(baseURL)
}

// Get weather for a city
func (service *Service) Get(city string) (Response, error) {
	var response Response
	if service.appID == "" {
		return response, errors.New("App ID not set")
	}
	params := map[string]string{
		"APPID": service.appID,
		"q":     city,
	}
	body, err := service.client.Get("weather", params)
	if err != nil {
		return response, err
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
