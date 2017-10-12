package usecase

import "wu/models"

type (
	AstronomyInteractor interface {
		ShowAstronomy(string) error
	}

	AstronomyPresenter interface {
		Print(*models.AstronomyDay) error
	}

	astronomyInteractor struct {
		repo      models.AstronomyRepository
		presenter AstronomyPresenter
	}
)

func NewAstronomyInteractor(repo models.AstronomyRepository, presenter AstronomyPresenter) AstronomyInteractor {
	return &astronomyInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (ai *astronomyInteractor) ShowAstronomy(location string) error {
	day, err := ai.repo.GetAstronomy(location)
	if err != nil {
		return err
	}

	return ai.presenter.Print(day)
}
