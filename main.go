package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

func getWeatherData(city string, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city>")
		return
	}

	city := os.Args[1]
	apiKey := "..." // Replace this with your API key

	weatherData, err := getWeatherData(city, apiKey)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}

	fmt.Printf("Weather in %s: %s\n", weatherData.Name, weatherData.Weather[0].Description)
	fmt.Printf("Temperature: %.1f°C\n", weatherData.Main.Temperature)
}
