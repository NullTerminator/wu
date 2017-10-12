package adapters

import "wu/models"

type (
	ForecastPresenter interface {
		Print([]*models.ForecastDay) error
	}
)
