package models

type (
	ForecastRepository interface {
		Get() ([]*ForecastDay, error)
	}

	ForecastDay struct {
		ForecastDate
		High           string
		Low            string
		Conditions     string
		ChanceOfPrecip int
		Rain           float32
		Snow           float32
	}

	ForecastDate struct {
		Pretty         string
		Day            int
		Month          int
		Year           int
		MonthName      string
		MonthNameShort string
		Weekday        string
		WeekdayShort   string
	}
)
