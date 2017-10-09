package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"wu/models"
)

const (
	API_ROOT = "http://api.wunderground.com/api/"
)

type (
	HttpHandler interface {
		Get(string) ([]byte, error)
	}

	wundergroundAdapter struct {
		key         string
		location    string
		httpHandler HttpHandler
	}

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
		Snow           Depth `json:'snow_allday"`
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

	Response struct {
		Error ForecastError
	}

	ForecastError struct {
		Type        string
		Description string
	}
)

func NewWundergroundAdapter(key, location string, handler HttpHandler) *wundergroundAdapter {
	return &wundergroundAdapter{
		key:         key,
		location:    location,
		httpHandler: handler,
	}
}

func (wa *wundergroundAdapter) Get() ([]*models.ForecastDay, error) {
	url := fmt.Sprintf("%s%s/forecast10day/q/%s.json", API_ROOT, wa.key, wa.location)
	body, err := wa.httpHandler.Get(url)
	if err != nil {
		return nil, err
	}

	var resp ForecastResponse
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if resp.Response.Error.Type != "" {
		return nil, errors.New(fmt.Sprintf("%s - %s", resp.Response.Error.Type, resp.Response.Error.Description))
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
