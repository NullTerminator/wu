package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/buger/goterm"
)

const (
	API_ROOT = "http://api.wunderground.com/api/"
)

type (
	ForecastConditions struct {
		Forecast Forecast
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
		year            int
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

func main() {
	key := os.Getenv("WU_KEY")
	url := fmt.Sprintf("%s%s/forecast10day/q/NH/Freedom.json", API_ROOT, key)

	res, err := http.Get(url)
	CheckError(err)
	if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad HTTP Status: %d\n", res.StatusCode)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var conds ForecastConditions
	jsonErr := json.Unmarshal(b, &conds)
	CheckError(jsonErr)

	goterm.Clear()
	var boxes []*goterm.Box

	for _, day := range conds.Forecast.Simpleforecast.Days {
		fmt.Println(day)
		box := goterm.NewBox(10|goterm.PCT, 20, 0)
		fmt.Fprintf(box, "%s\n%s / %s\n%.2f in\n%d%%", day.Date.Weekday_short, day.High.F, day.Low.F, day.Rain.In, day.ChanceOfPrecip)
		boxes = append(boxes, box)
	}

	for i, box := range boxes {
		goterm.Print(goterm.MoveTo(box.String(), i*10|goterm.PCT, 90|goterm.PCT))
	}
	goterm.Flush()
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error\n%v\n", err)
		os.Exit(1)
	}
}
