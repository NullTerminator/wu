package adapters

import (
	"fmt"
	"wu/models"

	"github.com/buger/goterm"
)

type (
	BoxAstronomyPresenter AstronomyPresenter

	boxAstronomyPresenter struct {
		logger Logger
	}
)

func NewBoxAstronomyPresenter(logger Logger) BoxAstronomyPresenter {
	return &boxAstronomyPresenter{
		logger: logger,
	}
}

func (presenter *boxAstronomyPresenter) Print(day *models.AstronomyDay) error {
	presenter.logger.Debug("Printing astronomy day")
	goterm.Clear()

	box := goterm.NewBox(20, 5, 0)
	fmt.Fprintf(box, "Sunrise: %s\nSunset: %s",
		day.Sunrise,
		day.Sunset)
	goterm.Print(goterm.MoveTo(box.String(), 40|goterm.PCT, 1))

	goterm.Flush()

	return nil
}
