package adapters

import (
	"encoding/json"
	"fmt"
	"wu/models"
)

type (
	AstronomyBody struct {
		Response Response
		SunPhase SunPhase `json:"sun_phase"`
	}

	SunPhase struct {
		Sunrise WuTime
		Sunset  WuTime
	}

	WuTime struct {
		Hour   string
		Minute string
	}
)

func (wa *wundergroundAdapter) GetAstronomy(location string) (*models.AstronomyDay, error) {
	wa.logger.Debugf("Getting astronomy for: %s", location)
	body, err := wa.get("astronomy", location)
	if err != nil {
		return nil, err
	}

	var resp AstronomyBody
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	err = wa.checkResponseErrors(resp.Response)
	if err != nil {
		return nil, err
	}

	day := &models.AstronomyDay{
		Sunrise: fmt.Sprintf("%s:%s", resp.SunPhase.Sunrise.Hour, resp.SunPhase.Sunrise.Minute),
		Sunset:  fmt.Sprintf("%s:%s", resp.SunPhase.Sunset.Hour, resp.SunPhase.Sunset.Minute),
	}

	return day, nil
}
