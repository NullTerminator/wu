package main

import (
	"os"
	"wu/adapters"
	"wu/infrastructure"
	"wu/usecase"
)

func main() {
	key := os.Getenv("WU_KEY")
	location := os.Getenv("WU_LOCATION")

	handler := &infrastructure.HttpHandler{}
	adapter := adapters.NewWundergroundAdapter(key, location, handler)
	presenter := &adapters.BoxPresenter{}

	forecaster := usecase.NewForecastInteractor(adapter, presenter)
	err := forecaster.Show()
	if err != nil {
		panic(err)
	}
}
