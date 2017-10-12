package adapters

import (
	"encoding/json"
	"wu/models"
)

type (
	ForecastResponse struct {
		Forecast Forecast
		Response Response
	}

	Forecast struct {
		Simpleforecast SimpleForecast
	}

	SimpleForecast struct {
		Days []Forecastday `json:"forecastday"`
	}

	Forecastday struct {
		Date           Forecastdate
		High           Temperature
		Low            Temperature
		Conditions     string
		ChanceOfPrecip int   `json:"pop"`
		Rain           Depth `json:"qpf_allday"`
		Snow           Depth `json:"snow_allday"`
	}

	Forecastdate struct {
		Pretty          string
		Day             int
		Month           int
		Year            int
		Monthname       string
		Monthname_short string
		Weekday         string
		Weekday_short   string
	}

	Temperature struct {
		F string `json:"Fahrenheit"`
	}

	Depth struct {
		In float32
	}
)

func (wa *wundergroundAdapter) GetTenDayForecast(location string) ([]*models.ForecastDay, error) {
	wa.logger.Debugf("Getting 10 day forecast for: %s", location)
	body, err := wa.get("forecast10day", location)
	if err != nil {
		return nil, err
	}

	var resp ForecastResponse
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	err = wa.checkResponseErrors(resp.Response)
	if err != nil {
		return nil, err
	}

	days := make([]*models.ForecastDay, len(resp.Forecast.Simpleforecast.Days))
	for i, wuday := range resp.Forecast.Simpleforecast.Days {
		day := models.ForecastDay{}
		day.Pretty = wuday.Date.Pretty
		day.Day = wuday.Date.Day
		day.Month = wuday.Date.Month
		day.Year = wuday.Date.Year
		day.MonthName = wuday.Date.Monthname
		day.MonthNameShort = wuday.Date.Monthname_short
		day.Weekday = wuday.Date.Weekday
		day.WeekdayShort = wuday.Date.Weekday_short
		day.High = wuday.High.F
		day.Low = wuday.Low.F
		day.Conditions = wuday.Conditions
		day.ChanceOfPrecip = wuday.ChanceOfPrecip
		day.Rain = wuday.Rain.In
		day.Snow = wuday.Snow.In
		days[i] = &day
	}

	return days, nil
}
