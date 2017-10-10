package adapters

import "wu/models"

type (
	Presenter interface {
		Print(days []*models.ForecastDay) error
	}
)
