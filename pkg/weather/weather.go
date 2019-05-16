package weather

import (
	"encoding/json"
	"log"

	"github.com/borzoj/go-lambda-test/pkg/http"
)

// Service weather api service
type Service struct {
	client *http.Client
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
func NewService(client *http.Client) (*Service, error) {
	return &Service{client: client}, nil
}

// Get weather for a city
func (service *Service) Get(city string) (Response, error) {
	var response Response
	params := map[string]string{
		"APPID": "50819be5818fbd89e2833c786d0a503e",
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
