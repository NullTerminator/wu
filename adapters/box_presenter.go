package adapters

import (
	"fmt"
	"wu/models"

	"github.com/buger/goterm"
)

type (
	BoxPresenter struct {
	}
)

func (presenter *BoxPresenter) Print(days []*models.ForecastDay) error {
	goterm.Clear()

	for i, day := range days {
		box := goterm.NewBox(10|goterm.PCT, 7, 0)
		fmt.Fprintf(box, "%s\n%s / %s\n%s\n%.2f in\n%d%%",
			day.WeekdayShort,
			day.High, day.Low,
			day.Conditions,
			day.Rain,
			day.ChanceOfPrecip)

		goterm.Print(goterm.MoveTo(box.String(), i*10|goterm.PCT, 1))
	}

	goterm.Flush()

	return nil
}
