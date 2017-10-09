package main

import (
	"flag"
	"os"
	"wu/adapters"
	"wu/infrastructure"
	"wu/usecase"
)

func main() {
	var (
		key      string
		location string
	)

	flag.StringVar(&key, "api_key", os.Getenv("WU_KEY"), "WeatherUnderground API Key")
	flag.StringVar(&location, "location", os.Getenv("WU_LOCATION"), "City and state location: (Freedom, NH)")
	flag.Parse()

	if key == "" {
		panic("API key required via WU_KEY env var or '--api_key' argument")
	}
	if location == "" {
		panic("location required via WU_LOCATION env var or '--location' argument")
	}

	handler := &infrastructure.HttpHandler{}
	adapter := adapters.NewWundergroundAdapter(key, location, handler)
	presenter := &adapters.BoxPresenter{}

	forecaster := usecase.NewForecastInteractor(adapter, presenter)
	err := forecaster.Show()
	if err != nil {
		panic(err)
	}
}
