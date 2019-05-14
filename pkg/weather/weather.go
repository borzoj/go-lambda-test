package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather description
type Weather struct {
	Description string `json:"description" binding:"required"`
	Main        string `json:"main" binding:"required"`
}

// Response api description
type Response struct {
	Weather []Weather `json:"weather" binding:"required"`
}

// Get weather for a city
func Get(city string) (Response, error) {
	var response Response

	owReq, _ := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	// if you appending to existing query this works fine
	q := owReq.URL.Query()
	q.Add("APPID", "50819be5818fbd89e2833c786d0a503e")
	q.Add("q", city)
	owReq.URL.RawQuery = q.Encode()

	client := &http.Client{}
	owResponse, err := client.Do(owReq)

	if err != nil {
		return response, err
	}
	defer owResponse.Body.Close()
	body, err := ioutil.ReadAll(owResponse.Body)
	log.Println(string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
