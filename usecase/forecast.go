package usecase

import "wu/models"

type (
	ForecastInteractor interface {
		Show() error
	}

	ForecastPresenter interface {
		Print([]*models.ForecastDay) error
	}

	forecastInteractor struct {
		Repo      models.ForecastRepository
		Presenter ForecastPresenter
	}
)

func NewForecastInteractor(repo models.ForecastRepository, presenter ForecastPresenter) ForecastInteractor {
	return &forecastInteractor{
		Repo:      repo,
		Presenter: presenter,
	}
}

func (fi *forecastInteractor) Show() error {
	days, err := fi.Repo.Get()
	if err != nil {
		return err
	}

	return fi.Presenter.Print(days)
}
