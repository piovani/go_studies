package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endpoint                = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByCity = "/weather?q=%s&appid=%s&units=metric"
)

type provider struct {
	apiKey string
}

func NewProvider(apiKey string) *provider {
	return &provider{
		apiKey: apiKey,
	}
}

func (p *provider) GetWeatherByCity(city string) (Wather, error) {
	var wather Wather
	path := fmt.Sprintf(pathFormatWeatherByCity, city, p.apiKey)
	u := endpoint + path

	res, err := http.Get(u)
	if err != nil {
		return wather, fmt.Errorf("openweather.GetWeatherByCity failed http GET: %s", err)
	}
	defer res.Body.Close()

	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return wather, fmt.Errorf("openweather.GetWeatherByCity failed reading body: %s", err)
	}

	var weatherRes WeatherReponse
	if err = json.Unmarshal(bodyRaw, &weatherRes); err != nil {
		return wather, fmt.Errorf("openweather.GetWeatherByCity failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return wather, fmt.Errorf("openweather.GetWeatherByCity got error from OpenWeather: %s", weatherRes.Message)
	}

	return weatherRes.ToWeather(), nil
}
