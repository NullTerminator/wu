package main

import (
	"flag"
	"os"
	"wu/adapters"
	"wu/infrastructure"
	"wu/logger"
	"wu/usecase"
)

func main() {
	key := flag.String("api_key", os.Getenv("WU_KEY"), "WeatherUnderground API Key")
	location := flag.String("location", os.Getenv("WU_LOCATION"), "City and state location: (Freedom, NH)")
	debug := flag.Bool("debug", false, "Debug logging")
	flag.BoolVar(debug, "d", false, "Debug logging")
	flag.Parse()

	if *key == "" {
		panic("API key required via WU_KEY env var or '--api_key' argument")
	}
	if *location == "" {
		panic("location required via WU_LOCATION env var or '--location' argument")
	}
	if *debug {
		logger.SetLevel(logger.DEBUG)
	} else {
		logger.SetLevel(logger.INFO)
	}

	handler := infrastructure.NewHttpHandler(logger.Logger)
	adapter := adapters.NewWundergroundAdapter(*key, handler, logger.Logger)
	presenter := adapters.NewBoxPresenter(logger.Logger)

	forecaster := usecase.NewForecastInteractor(adapter, presenter)
	err := forecaster.Show(*location)
	if err != nil {
		panic(err)
	}
}
