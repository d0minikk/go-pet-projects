package openmeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"context"
	"time"
)

type WeatherResponse struct {
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	GenerationTimeMs    float64        `json:"generationtime_ms"`
	UtcOffsetSeconds    int            `json:"utc_offset_seconds"`
	Timezone            string         `json:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbreviation"`
	Elevation           float64        `json:"elevation"`
	CurrentUnits        CurrentUnits   `json:"current_units"`
	Current             CurrentWeather `json:"current"`
}

type CurrentUnits struct {
	Time            string `json:"time"`
	Interval        string `json:"interval"`
	Temperature2M   string `json:"temperature_2m"`
	WindSpeed10M    string `json:"wind_speed_10m"`
}

type CurrentWeather struct {
	Time          string  `json:"time"`
	Interval      int     `json:"interval"`
	Temperature2M float64 `json:"temperature_2m"`
	WindSpeed10M  float64 `json:"wind_speed_10m"`
}

const apiEndpoint = "https://api.open-meteo.com/v1/forecast"

func GetCurrentWeather(ctx context.Context, latitude, longitude float64) (*WeatherResponse, error) {
	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%f", latitude))
	params.Add("longitude", fmt.Sprintf("%f", longitude))
	params.Add("current", "temperature_2m,wind_speed_10m")

	requestURL := fmt.Sprintf("%s?%s", apiEndpoint, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("User-Agent", "weather-app/1.0")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var weatherResp WeatherResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &weatherResp, nil
}

