package main

type Wather struct {
	Temp     float32
	Pressure float32
	MinTemp  float32
	MaxTemp  float32
}

type WeatherProvider interface {
	GetWeatherByCity(city string) (Wather, error)
}