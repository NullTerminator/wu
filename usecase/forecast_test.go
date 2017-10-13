package usecase_test

import (
	"errors"
	"testing"

	"wu/models"
	mmocks "wu/models/mocks"
	"wu/usecase"
	ucmocks "wu/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestShowTenDay(t *testing.T) {
	location := "foo"
	days := []*models.ForecastDay{}

	mockRepo := new(mmocks.ForecastRepository)
	mockRepo.On("GetTenDayForecast", location).Return(days, nil)

	mockPresenter := new(ucmocks.ForecastPresenter)
	mockPresenter.On("Print", days).Return(nil)

	actor := usecase.NewForecastInteractor(mockRepo, mockPresenter)
	err := actor.ShowTenDay(location)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestShowTenDayGetsErrorFetchingDays(t *testing.T) {
	location := "foo"
	errMsg := "bad forecast request"
	err := errors.New(errMsg)

	mockRepo := new(mmocks.ForecastRepository)
	mockRepo.On("GetTenDayForecast", location).Return(nil, err)

	mockPresenter := new(ucmocks.ForecastPresenter)

	actor := usecase.NewForecastInteractor(mockRepo, mockPresenter)
	err = actor.ShowTenDay(location)
	if assert.NotNil(t, err) {
		assert.Equal(t, errMsg, err.Error())
	}
}
