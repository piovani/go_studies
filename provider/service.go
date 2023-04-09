package main

import "fmt"

type Service struct {
	WeatherProvider WeatherProvider
}

func NewService(p WeatherProvider) *Service {
	return &Service{
		WeatherProvider: p,
	}
}

func (s *Service) WhatToWear(city string) (string, error) {
	w, err := s.WeatherProvider.GetWeatherByCity(city)
	if err != nil {
		return "", fmt.Errorf("WhatToWear: %w", err)
	}

	if w.Temp < 21 {
		return "long sleeves", nil
	}

	return "short sleeves", nil
}
