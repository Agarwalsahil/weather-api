package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"weather-api/graph/model"
)

const (
	weatherApiUrl = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"
	unit          = "metric"
)

func GetWeatherFromAPI(city string, givenDate string) (*model.WeatherResponse, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY environment variable not set")
	}

	url := fmt.Sprintf("%s/%s/%s?unitGroup=%s&key=%s", weatherApiUrl, city, givenDate, unit, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to send the request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, errRead := io.ReadAll(resp.Body)

	if errRead != nil {
		return nil, fmt.Errorf("error in reading the response: %w", errRead)
	}

	var weather model.WeatherResponse
	if errMarshal := json.Unmarshal(body, &weather); errMarshal != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", errMarshal)
	}

	return &weather, nil
}
