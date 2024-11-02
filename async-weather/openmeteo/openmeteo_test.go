package openmeteo

import (
	"context"
	"testing"
)

func TestGetCurrentWeather(t *testing.T) {
	ctx := context.Background()
	lat := 52.52
	lng := 13.41

	weatherData, err := GetCurrentWeather(ctx, lat, lng)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if weatherData.Current.Temperature2M == 0 {
		t.Errorf("Expected temperature to be non-zero")
	}

	if weatherData.Current.WindSpeed10M == 0 {
		t.Errorf("Expected wind speed to be non-zero")
	}
}