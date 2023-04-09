package main

type WeatherReponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	}
}

func (r WeatherReponse) ToWeather() Wather {
	return Wather{
		Temp:     r.Main.Temp,
		Pressure: r.Main.Pressure,
		MinTemp:  r.Main.TempMin,
		MaxTemp:  r.Main.TempMax,
	}
}
