package adapters

import (
	"errors"
	"fmt"
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
		httpHandler HttpHandler
		logger      Logger
	}

	Response struct {
		Error WunderError
	}

	WunderError struct {
		Type        string
		Description string
	}
)

func NewWundergroundAdapter(key string, handler HttpHandler, logger Logger) *wundergroundAdapter {
	return &wundergroundAdapter{
		key:         key,
		httpHandler: handler,
		logger:      logger,
	}
}

func (wa *wundergroundAdapter) get(path string, location string) ([]byte, error) {
	url := fmt.Sprintf("%s%s/%s/q/%s.json", API_ROOT, wa.key, path, location)
	return wa.httpHandler.Get(url)
}

func (wa *wundergroundAdapter) checkResponseErrors(response Response) error {
	if response.Error.Type != "" {
		return errors.New(fmt.Sprintf("%s - %s", response.Error.Type, response.Error.Description))
	}

	return nil
}
